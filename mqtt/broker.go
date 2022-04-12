package mqtt

import (
	"encoding/binary"
	"github.com/google/uuid"
	"github.com/zgwit/iot-master/events"
	"github.com/zgwit/iot-master/log"
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
	events.EventEmitter

	//Subscribe tree
	subTree SubTree

	//Retain tree
	retainTree RetainTree

	//ClientId->Client
	clients sync.Map // map[string]*Client
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
	client := NewClient(conn)
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
		for _, pkt := range ms {
			h.handle(pkt, client)
		}
	}

	_ = client.Close()
}

func (h *Broker) Receive2(conn net.Conn) {
	//TODO 先解析第一个包，而且必须是Connect
	client := NewClient(conn)

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
		pkt, err := DecodePacket(buf[:packLen])
		if err != nil {
			log.Println(err)
			break
		}

		//处理消息
		//TODO 可以放入队列
		h.handle(pkt, client)

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

	_ = client.Close()
}

func (h *Broker) handle(pkt Packet, client *Client) {
	switch pkt.Type() {
	case CONNECT:
		h.handleConnect(pkt.(*Connect), client)
	case PUBLISH:
		h.handlePublish(pkt.(*Publish), client)
	case PUBACK:
		client.pub1.Delete(pkt.(*PubAck).PacketId())
	case PUBREC:
		pkt.SetType(PUBREL)
		client.dispatch(pkt)
	case PUBREL:
		pkt.SetType(PUBCOMP)
		client.dispatch(pkt)
	case PUBCOMP:
		client.pub2.Delete(pkt.(*PubComp).PacketId())
	case SUBSCRIBE:
		h.handleSubscribe(pkt.(*Subscribe), client)
	case UNSUBSCRIBE:
		h.handleUnSubscribe(pkt.(*UnSubscribe), client)
	case PINGREQ:
		pkt.SetType(PINGRESP)
		client.dispatch(pkt)
	case DISCONNECT:
		h.handleDisconnect(pkt.(*DisConnect), client)
	}
}

func (h *Broker) handleConnect(pkt *Connect, client *Client) {
	ack := CONNACK.NewPacket().(*Connack)

	//TODO 验证用户名密码
	//if h.onConnect != nil {
	//	if !h.onConnect(pkt, client) {
	//		ack.SetCode(CONNACK_INVALID_USERNAME_PASSWORD)
	//		client.dispatch(ack)
	//		// 断开
	//		_ = client.Close()
	//		return
	//	}
	//}

	h.Emit("connect", pkt, client)

	var clientId string
	if len(pkt.ClientId()) == 0 {

		if !pkt.CleanSession() {
			//TODO 无ID，必须是清空会话 error
			//return
			_ = client.Close()
			return
		}

		// Generate unique clientId (uuid random)
		clientId = uuid.New().String()
		//UUID不用验重了
		//for { if _, ok := h.clients.Load(clientId); !ok { break } }
	} else {
		clientId = string(pkt.ClientId())

		if v, ok := h.clients.Load(clientId); ok {
			b := v.(*Client)
			// ClientId is already used
			if !b.closed {
				//error reject
				ack.SetCode(CONNACK_UNAVAILABLE)
				client.dispatch(ack)
				_ = client.Close()
				return
			} else {
				if !pkt.CleanSession() {
					//TODO 复制内容
					client.keepAlive = b.keepAlive
					client.will = b.will
					//client.pub1 = b.pub1 //sync.Map不能直接复制。。。。
					b.pub1.Range(func(key, value interface{}) bool {
						client.pub1.Store(key, value)
						return true
					})
					//client.pub2 = b.pub2
					b.pub2.Range(func(key, value interface{}) bool {
						client.pub2.Store(key, value)
						return true
					})
					//client.receivedPub2 = b.receivedPub2
					b.receivedPub2.Range(func(key, value interface{}) bool {
						client.receivedPub2.Store(key, value)
						return true
					})
					client.packetId = b.packetId

					//ack.SetSessionPresent(true)
				}
			}
		}

		h.clients.Store(clientId, client)
	}

	client.clientId = clientId

	if pkt.KeepAlive() > 0 {
		client.timeout = time.Second * time.Duration(pkt.KeepAlive()) * 3 / 2
	}

	//TODO 如果发生错误，与客户端断开连接
	ack.SetCode(CONNACK_ACCEPTED)
	client.dispatch(ack)
}

func (h *Broker) handlePublish(pkt *Publish, client *Client) {

	h.Emit("publish", pkt, client)

	qos := pkt.Qos()
	if qos == Qos0 {
		//不需要回复puback
	} else if qos == Qos1 {
		//Reply PUBACK
		ack := PUBACK.NewPacket().(*PubAck)
		ack.SetPacketId(pkt.PacketId())
		client.dispatch(ack)
	} else if qos == Qos2 {
		//Save & Send PUBREC
		client.receivedPub2.Store(pkt.PacketId(), pkt)
		ack := PUBREC.NewPacket().(*PubRec)
		ack.SetPacketId(pkt.PacketId())
		client.dispatch(ack)
	} else {
		//TODO error

	}

	if err := ValidTopic(pkt.Topic()); err != nil {
		log.Println("Topic invalid ", err)
		return
	}

	if pkt.Retain() {
		if len(pkt.Payload()) == 0 {
			h.retainTree.UnRetain(client.clientId)
		} else {
			h.retainTree.Retain(pkt.Topic(), client.clientId, pkt)
		}
	}

	//Fetch subscribers
	subs := make(map[string]MsgQos)
	h.subTree.Publish(pkt.Topic(), subs)

	//Send publish message
	for clientId, qos := range subs {
		if b, ok := h.clients.Load(clientId); ok {
			bb := b.(*Client)
			if bb.closed {
				continue
			}

			//clone new pub
			pub := *pkt
			pub.SetRetain(false)
			if pkt.Qos() > qos {
				pub.SetQos(qos)
			}
			bb.dispatch(&pub)
		}
	}
}

func (h *Broker) handleSubscribe(pkt *Subscribe, client *Client) {
	ack := SUBACK.NewPacket().(*SubAck)
	ack.SetPacketId(pkt.PacketId())

	//TODO 外部验证
	//if h.onSubscribe != nil {
	//	if !h.onSubscribe(pkt, client) {
	//		//回复失败
	//		ack.AddCode(SUB_CODE_ERR)
	//		client.dispatch(ack)
	//		return
	//	}
	//}
	h.Emit("subscribe", pkt, client)

	for _, st := range pkt.Topics() {
		//log.Print("Subscribe ", string(st.Topic()))
		if err := ValidSubscribe(st.Topic()); err != nil {
			log.Println("Invalid topic ", err)
			//log error
			ack.AddCode(SUB_CODE_ERR)
		} else {
			h.subTree.Subscribe(st.Topic(), client.clientId, st.Qos())

			ack.AddCode(SubCode(st.Qos()))
			h.retainTree.Fetch(st.Topic(), func(clientId string, pub *Publish) {
				//clone new pub
				p := *pub
				p.SetRetain(true)
				if pkt.Qos() > st.Qos() {
					p.SetQos(st.Qos())
				}
				client.dispatch(&p)
			})
		}
	}
	client.dispatch(ack)
}

func (h *Broker) handleUnSubscribe(pkt *UnSubscribe, client *Client) {
	//外部验证
	//if h.onUnSubscribe != nil {
	//	h.onUnSubscribe(pkt, client)
	//}
	h.Emit("unsubscribe", pkt, client)

	ack := UNSUBACK.NewPacket().(*UnSubAck)
	for _, t := range pkt.Topics() {
		//log.Print("UnSubscribe ", string(t))
		if err := ValidSubscribe(t); err != nil {
			log.Error(err)
		} else {
			h.subTree.UnSubscribe(t, client.clientId)
		}
	}
	client.dispatch(ack)
}

func (h *Broker) handleDisconnect(pkt *DisConnect, client *Client) {
	//if h.onDisconnect != nil {
	//	h.onDisconnect(pkt, client)
	//}
	h.Emit("disconnect", pkt, client)

	h.clients.Delete(client.clientId)
	_ = client.Close()
}
