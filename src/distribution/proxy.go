package distribution

import (
	infra "github.com/ruybrito106/mom-client-go/src/infrastructure"
)

const (
	PublisherType  = "pub"
	SubscriberType = "sub"
)

type Proxy struct {
	// manager *CallManager
	addr string
}

func New(addr string) (*Proxy, error) {

	return &Proxy{
		// manager: &CallManager{
		// 	make(map[string](func() interface{}), 0),
		// },
		addr: addr,
	}, nil

}

func (p *Proxy) Publish(topic, msg string) error {

	conn := &infra.Connection{}
	err := conn.OpenTCP(p.addr)
	if err != nil {
		return err
	}

	packet := &Packet{
		Type: "pub",
		Message: &Message{
			Topic:   topic,
			Content: msg,
		},
	}

	return conn.Write(packet)

}

func (p *Proxy) Subscribe(topic string, call func(interface{})) error {

	conn := &infra.Connection{}
	err := conn.OpenTCP(p.addr)
	if err != nil {
		return err
	}

	packet := &Packet{
		Type: "sub",
		Message: &Message{
			Topic: topic,
		},
	}

	conn.Write(packet)

	conn.Reader(call)

	return nil

}
