package service

import "net"

type Link interface {
	Write(data []byte) error
	Read(data []byte) (int, error)
	Close() error
}

type NetLink struct {
	Id   int
	conn net.Conn
}

func NewNetLink(conn net.Conn) *NetLink {
	return &NetLink{conn: conn}
}

func (l *NetLink) Write(data []byte) error {
	_, err := l.conn.Write(data)
	return err
}

func (l *NetLink) Read(data []byte) (int, error) {
	return l.conn.Read(data)
}

func (l *NetLink) Close() error {
	return l.conn.Close()
}
