package service

import "net"

type TcpClient struct {
	Addr string
	conn net.Conn
	link Link
}

func (c *TcpClient) Open() error {
	conn, err := net.Dial("tcp", c.Addr)
	if err != nil {
		return err
	}
	c.conn = conn
	//TODO 封装conn
	return nil
}

func (c *TcpClient) HasAcceptor() bool {
	return false
}

func (c *TcpClient) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil //TODO return error
}

func (c *TcpClient) GetLink(id int) (Link, error) {
	return c.link, nil
}
