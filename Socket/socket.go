package socket

import (
	"net"
)

type Network struct {
	Con      net.Conn
	IsClosed bool
}

func Connect() *Network {
	// 连接到服务器
	conn, err := net.Dial("tcp", "localhost:8081")
	if err != nil {
		return nil
	}

	net := Network{Con: conn, IsClosed: false}

	conn.Write([]byte("Client connet"))
	return &net
}

func (n *Network) SendMessage(s string) error {
	_, err := n.Con.Write([]byte(s))
	if err != nil {
		return err
	}
	return nil
}

// Close 方法用于关闭连接
func (n *Network) Close() error {
	if n.Con != nil && !n.IsClosed {
		err := n.Con.Close()
		if err != nil {
			return err
		}
		n.IsClosed = true
	}
	return nil
}
