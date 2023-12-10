package socket

import (
	"net"
)

const (
	socketLocalHost = "localhost:8081"
)

type Network struct {
	Con      net.Conn
	IsClosed bool
}

func Connect() *Network {
	// 連線至伺服器
	conn, err := net.Dial("tcp", socketLocalHost)
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
