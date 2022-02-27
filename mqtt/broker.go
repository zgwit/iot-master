package mqtt

import (
	"encoding/binary"
	"github.com/google/uuid"
	"log"
	"net"
	"sync"
	"time"
)

func reAlloc(buf []byte, l int) []byte {
	b := make([]byte, l)
	copy(b, buf)
	return b
}

type Broker struct {

	//Subscribe tree
	subTree SubTree

	//Retain tree
	retainTree RetainTree

	//ClientId->Client
	bees sync.Map // map[string]*Client

	onConnect     func(*Connect, *Client) bool
	onPublish     func(*Publish, *Client) bool
	onSubscribe   func(*Subscribe, *Client) bool
	onUnSubscribe func(*UnSubscribe, *Client)
	onDisconnect  func(*DisConnect, *Client)
}

//TODO 添加参数
func NewBroker() *Broker {
	return &Broker{}
}

func (h *Broker) ListenAndServe(addr string) error {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	go h.Serve(ln)
	return nil
}

func (h *Broker) Serve(ln net.Listener) {
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			break
		}

		//process
		go h.Receive(conn)
	}
}

func (h *Broker) Receive(conn net.Conn) {
	//TODO 先解析第一个包，而且必须是Connect
	bee := NewBee(conn)
	var parser Parser

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Println(err)
			break
		}

		ms := parser.Parse(buf[:n])

		//处理消息
		//TODO 可以放入队列
		for _, msg := range ms {
			h.handle(msg, bee)
		}
	}

	_ = bee.Close()
}

func (h *Broker) Receive2(conn net.Conn) {
	//TODO 先解析第一个包，而且必须是Connect
	bee := NewBee(conn)

	bufSize := 6
	buf := make([]byte, bufSize)
	of := 0
	for {
		//TODO 先解析
		n, err := conn.Read(buf[of:])
		if err != nil {
			log.Println(err)
			break
		}
		ln := of + n

		if ln < 2 {
			of = ln
			continue
		}

		//解析包头，包体

		//读取Remain Length
		rl, rll := binary.Uvarint(buf[1:])
		remainLength := int(rl)
		packLen := remainLength + rll + 1

		//读取未读完的包体
		if packLen > bufSize {
			buf = reAlloc(buf, packLen)

			//直至将全部包体读完
			o := ln
			for o < packLen {
				n, err = conn.Read(buf[o:])
				if err != nil {
					log.Println(err)
					//return
					break
				}
				o += n
			}
			//一般不会发生
			if o < packLen {
				break
			}
		}

		//解析消息
		msg, err := Decode(buf[:packLen])
		if err != nil {
			log.Println(err)
			break
		}

		//处理消息
		//TODO 可以放入队列
		h.handle(msg, bee)

		//解析 剩余内容
		if packLen < bufSize {
			buf = reAlloc(buf[packLen:], bufSize-packLen)
			of = bufSize - packLen
			//TODO 剩余内容可能已经包含了消息，不用再Read，直接解析
		} else {
			buf = make([]byte, bufSize)
			of = 0
		}
	}

	_ = bee.Close()
}

func (h *Broker) handle(msg Packet, bee *Client) {
	switch msg.Type() {
	case CONNECT:
		h.handleConnect(msg.(*Connect), bee)
	case PUBLISH:
		h.handlePublish(msg.(*Publish), bee)
	case PUBACK:
		bee.pub1.Delete(msg.(*PubAck).PacketId())
	case PUBREC:
		msg.SetType(PUBREL)
		bee.dispatch(msg)
	case PUBREL:
		msg.SetType(PUBCOMP)
		bee.dispatch(msg)
	case PUBCOMP:
		bee.pub2.Delete(msg.(*PubComp).PacketId())
	case SUBSCRIBE:
		h.handleSubscribe(msg.(*Subscribe), bee)
	case UNSUBSCRIBE:
		h.handleUnSubscribe(msg.(*UnSubscribe), bee)
	case PINGREQ:
		msg.SetType(PINGRESP)
		bee.dispatch(msg)
	case DISCONNECT:
		h.handleDisconnect(msg.(*DisConnect), bee)
	}
}

func (h *Broker) handleConnect(msg *Connect, bee *Client) {
	ack := CONNACK.NewPacket().(*Connack)

	//验证用户名密码
	if h.onConnect != nil {
		if !h.onConnect(msg, bee) {
			ack.SetCode(CONNACK_INVALID_USERNAME_PASSWORD)
			bee.dispatch(ack)
			// 断开
			_ = bee.Close()
			return
		}
	}

	var clientId string
	if len(msg.ClientId()) == 0 {

		if !msg.CleanSession() {
			//TODO 无ID，必须是清空会话 error
			//return
			_ = bee.Close()
			return
		}

		// Generate unique clientId (uuid random)
		clientId = uuid.New().String()
		//UUID不用验重了
		//for { if _, ok := h.bees.Load(clientId); !ok { break } }
	} else {
		clientId = string(msg.ClientId())

		if v, ok := h.bees.Load(clientId); ok {
			b := v.(*Client)
			// ClientId is already used
			if !b.closed {
				//error reject
				ack.SetCode(CONNACK_UNAVAILABLE)
				bee.dispatch(ack)
				_ = bee.Close()
				return
			} else {
				if !msg.CleanSession() {
					//TODO 复制内容
					bee.keepAlive = b.keepAlive
					bee.will = b.will
					bee.pub1 = b.pub1 //sync.Map不能直接复制。。。。
					bee.pub2 = b.pub2
					bee.recvPub2 = b.recvPub2
					bee.packetId = b.packetId

					//ack.SetSessionPresent(true)
				}
			}
		}

		h.bees.Store(clientId, bee)
	}

	bee.clientId = clientId

	if msg.KeepAlive() > 0 {
		bee.timeout = time.Second * time.Duration(msg.KeepAlive()) * 3 / 2
	}

	//TODO 如果发生错误，与客户端断开连接
	ack.SetCode(CONNACK_ACCEPTED)
	bee.dispatch(ack)
}

func (h *Broker) handlePublish(msg *Publish, bee *Client) {
	//外部验证
	if h.onPublish != nil {
		if !h.onPublish(msg, bee) {
			return
		}
	}

	qos := msg.Qos()
	if qos == Qos0 {
		//不需要回复puback
	} else if qos == Qos1 {
		//Reply PUBACK
		ack := PUBACK.NewPacket().(*PubAck)
		ack.SetPacketId(msg.PacketId())
		bee.dispatch(ack)
	} else if qos == Qos2 {
		//Save & Send PUBREC
		bee.recvPub2.Store(msg.PacketId(), msg)
		ack := PUBREC.NewPacket().(*PubRec)
		ack.SetPacketId(msg.PacketId())
		bee.dispatch(ack)
	} else {
		//TODO error

	}

	if err := ValidTopic(msg.Topic()); err != nil {
		//TODO log
		log.Println("Topic invalid ", err)
		return
	}

	if msg.Retain() {
		if len(msg.Payload()) == 0 {
			h.retainTree.UnRetain(bee.clientId)
		} else {
			h.retainTree.Retain(msg.Topic(), bee.clientId, msg)
		}
	}

	//Fetch subscribers
	subs := make(map[string]MsgQos)
	h.subTree.Publish(msg.Topic(), subs)

	//Send publish message
	for clientId, qos := range subs {
		if b, ok := h.bees.Load(clientId); ok {
			bb := b.(*Client)
			if bb.closed {
				continue
			}

			//clone new pub
			pub := *msg
			pub.SetRetain(false)
			if msg.Qos() > qos {
				pub.SetQos(qos)
			}
			bb.dispatch(&pub)
		}
	}
}

func (h *Broker) handleSubscribe(msg *Subscribe, bee *Client) {
	ack := SUBACK.NewPacket().(*SubAck)
	ack.SetPacketId(msg.PacketId())

	//外部验证
	if h.onSubscribe != nil {
		if !h.onSubscribe(msg, bee) {
			//回复失败
			ack.AddCode(SUB_CODE_ERR)
			bee.dispatch(ack)
			return
		}
	}

	for _, st := range msg.Topics() {
		//log.Print("Subscribe ", string(st.Topic()))
		if err := ValidSubscribe(st.Topic()); err != nil {
			log.Println("Invalid topic ", err)
			//log error
			ack.AddCode(SUB_CODE_ERR)
		} else {
			h.subTree.Subscribe(st.Topic(), bee.clientId, st.Qos())

			ack.AddCode(SubCode(st.Qos()))
			h.retainTree.Fetch(st.Topic(), func(clientId string, pub *Publish) {
				//clone new pub
				p := *pub
				p.SetRetain(true)
				if msg.Qos() > st.Qos() {
					p.SetQos(st.Qos())
				}
				bee.dispatch(&p)
			})
		}
	}
	bee.dispatch(ack)
}

func (h *Broker) handleUnSubscribe(msg *UnSubscribe, bee *Client) {
	//外部验证
	if h.onUnSubscribe != nil {
		h.onUnSubscribe(msg, bee)
	}

	ack := UNSUBACK.NewPacket().(*UnSubAck)
	for _, t := range msg.Topics() {
		//log.Print("UnSubscribe ", string(t))
		if err := ValidSubscribe(t); err != nil {
			//TODO log
			log.Println(err)
		} else {
			h.subTree.UnSubscribe(t, bee.clientId)
		}
	}
	bee.dispatch(ack)
}

func (h *Broker) handleDisconnect(msg *DisConnect, bee *Client) {
	if h.onDisconnect != nil {
		h.onDisconnect(msg, bee)
	}

	h.bees.Delete(bee.clientId)
	_ = bee.Close()
}

func (h *Broker) OnConnect(fn func(*Connect, *Client) bool) {
	h.onConnect = fn
}
func (h *Broker) OnPublish(fn func(*Publish, *Client) bool) {
	h.onPublish = fn
}
func (h *Broker) OnSubscribe(fn func(*Subscribe, *Client) bool) {
	h.onSubscribe = fn
}
func (h *Broker) OnUnSubscribe(fn func(*UnSubscribe, *Client)) {
	h.onUnSubscribe = fn
}
func (h *Broker) OnDisconnect(fn func(*DisConnect, *Client)) {
	h.onDisconnect = fn
}
