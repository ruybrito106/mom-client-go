package infrastructure

import (
	"net"
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"
)

const (
	transportProtocolTCP = "tcp"
)

type Connection struct {
	conn *websocket.Conn
}

func (c *Connection) OpenTCP(addr string) error {
	u, err := url.Parse("ws://" + addr)
	if err != nil {
		return err
	}

	tcpConn, err := net.Dial(transportProtocolTCP, u.Host)
	wsHeader := http.Header{
		"Origin": {"ws://" + addr},
	}

	wsConn, _, err := websocket.NewClient(tcpConn, u, wsHeader, 1024, 1024)
	if err != nil {
		return err
	}

	c.conn = wsConn

	return nil
}

func (c *Connection) Write(v interface{}) error {
	return c.conn.WriteJSON(v)
}

func (c *Connection) Read(callback func(interface{})) error {
	var v interface{}
	err := c.conn.ReadJSON(v)
	callback(v)
	return err
}

func (c *Connection) Reader(callback func(interface{})) {
	for {
		var v interface{}
		err := c.conn.ReadJSON(v)
		if err != nil {
			break
		}

		callback(v)
	}
	c.conn.Close()
}
