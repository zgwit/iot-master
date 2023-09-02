package vnet

import (
	"net"
	"time"
)

type vAddr struct {
}

func (a *vAddr) Network() string {
	return "vnet"
}

func (a *vAddr) String() string {
	return "vnet"
}

type VNet struct {
	peer *VNet
	buf  chan []byte
}

func (c *VNet) Close() error {
	close(c.buf)
	return nil
}

func (c *VNet) Read(b []byte) (int, error) {
	bb := <-c.buf
	n := copy(b, bb)
	//TODO 处理过长
	return n, nil
}

func (c *VNet) Write(b []byte) (int, error) {
	c.buf <- b
	return len(b), nil
}

func (c *VNet) LocalAddr() net.Addr                { return &vAddr{} }
func (c *VNet) RemoteAddr() net.Addr               { return &vAddr{} }
func (c *VNet) SetDeadline(t time.Time) error      { return nil }
func (c *VNet) SetReadDeadline(t time.Time) error  { return nil }
func (c *VNet) SetWriteDeadline(t time.Time) error { return nil }

func New() (*VNet, *VNet) {
	var c1, c2 VNet
	c1.peer = &c2
	c2.peer = &c1
	return &c1, &c2
}
