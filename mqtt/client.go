package mqtt

import (
	"log"
	"net"
	"sync"
	"time"
)

type Client struct {
	//Client ID (from CONNECT)
	clientId string

	//Keep Alive (from CONNECT)
	keepAlive int

	//will topic (from CONNECT)
	will *Publish

	//Qos1 Qos2
	pub1 sync.Map // map[uint16]*packet.Publish
	pub2 sync.Map // map[uint16]*packet.Publish

	//Received Qos2 Publish
	receivedPub2 sync.Map // map[uint16]*packet.Publish

	//Increment 0~65535
	packetId uint16

	conn net.Conn

	//消息发送队列，避免主协程任务过重
	pktQueue chan Packet

	timeout time.Duration

	//关闭标记
	closed bool
}

func NewClient(conn net.Conn) *Client {
	return &Client{
		conn: conn,
		//timeout: time.Hour * 24,
	}
}

func (b *Client) ClientId() string {
	return b.clientId
}

func (b *Client) Disconnect() error {
	_ = b.send(&DisConnect{})
	return b.Close()
}

func (b *Client) Close() error {
	err := b.conn.Close()
	b.closed = true
	return err
}

func (b *Client) send(pkt Packet) error {
	//log.Printf("Send message to %s: %s QOS(%d) DUP(%t) RETAIN(%t)", b.clientId, pkt.Type().Name(), pkt.Qos(), pkt.Dup(), pkt.Retain())
	if head, payload, err := pkt.Encode(); err != nil {
		return err
	} else {
		//err := b.conn.SetWriteDeadline(time.Now().Add(b.timeout))
		_, err = b.conn.Write(head)
		if err != nil {
			// 关闭client
			return err
		}
		if payload != nil && len(payload) > 0 {
			_, err = b.conn.Write(payload)
			if err != nil {
				// 关闭client
				return err
			}
		}
	}

	if pkt.Type() == PUBLISH {
		pub := pkt.(*Publish)
		//Publish Qos1 Qos2 Need store
		if pkt.Qos() == Qos1 {
			b.pub1.Store(pub.PacketId(), pub)
		} else if pkt.Qos() == Qos2 {
			b.pub2.Store(pub.PacketId(), pub)
		}
	}

	return nil
}

func (b *Client) dispatch(pkt Packet) {
	if b.pktQueue != nil {
		b.pktQueue <- pkt
		return
	}

	err := b.send(pkt)
	if err != nil {
		log.Println(err)
		//TODO 关闭
	}
}

func (b *Client) sender() {
	for b.conn != nil {
		//TODO select or close
		pkt := <-b.pktQueue
		err := b.send(pkt)
		if err != nil {
			break 
		}
	}
}
