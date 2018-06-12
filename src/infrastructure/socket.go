package infrastructure

import (
	"errors"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/eapache/go-resiliency/retrier"
	"github.com/gorilla/websocket"
)

const (
	transportProtocolTCP = "tcp"
)

type Connection struct {
	conn *websocket.Conn
}

type connectionClassifier struct{}

func (c connectionClassifier) Classify(err error) retrier.Action {
	if err == nil {
		return retrier.Succeed
	}
	return retrier.Retry
}

func (c *Connection) OpenTCP(addr string) error {
	u, err := url.Parse("ws://" + addr)
	if err != nil {
		return err
	}

	openConnectionRetrier := retrier.New(
		retrier.ExponentialBackoff(6, 100*time.Millisecond),
		connectionClassifier{},
	)

	err = openConnectionRetrier.Run(func() error {

		tcpConn, err := net.Dial(transportProtocolTCP, u.Host)
		wsHeader := http.Header{
			"Origin": {"ws://" + addr},
		}

		if err != nil {
			return err
		}

		wsConn, _, err := websocket.NewClient(tcpConn, u, wsHeader, 1024, 1024)
		if err != nil {
			return err
		}

		c.conn = wsConn

		return nil

	})

	return err

}

func (c *Connection) Write(v interface{}) error {
	return c.conn.WriteJSON(v)
}

func (c *Connection) readWithTimeout(d time.Duration) (v interface{}, err error) {

	successCh := make(chan bool)
	timeoutCh := time.After(d)

	go func() {
		err = c.conn.ReadJSON(&v)
		successCh <- true
	}()

	select {
	case <-successCh:
		return
	case <-timeoutCh:
		return v, errors.New("timeout")
	}

}

func (c *Connection) Read(callback func(interface{})) error {
	v, err := c.readWithTimeout(1 * time.Second)
	if err != nil {
		return err
	}

	callback(v)
	return nil
}

func (c *Connection) ReaderWithTimeout(callback func(interface{})) {
	for {
		v, err := c.readWithTimeout(200 * time.Millisecond)
		if err != nil {
			break
		}

		callback(v)
	}
}

func (c *Connection) Reader(callback func(interface{})) {
	for {
		var v interface{}
		err := c.conn.ReadJSON(&v)
		if err != nil {
			break
		}

		callback(v)
	}
}

func (c *Connection) Close() error {
	return c.conn.Close()
}
