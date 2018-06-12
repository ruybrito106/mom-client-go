package distribution

import (
	infra "github.com/ruybrito106/mom-client-go/src/infrastructure"
)

const (
	PublisherType  = "pub"
	SubscriberType = "sub"
)

type Proxy struct {
	addr string
}

func New(addr string) (*Proxy, error) {

	return &Proxy{
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
		Config: &Config{
			Type: "real-time",
		},
		Message: &Message{
			Topic: topic,
		},
	}

	conn.Write(packet)
	conn.Reader(call)

	return nil

}

func (p *Proxy) SubscribeForLatest(topic string, call func(interface{})) error {

	conn := &infra.Connection{}
	err := conn.OpenTCP(p.addr)
	if err != nil {
		return err
	}

	packet := &Packet{
		Type: "sub",
		Config: &Config{
			Type: "latest",
		},
		Message: &Message{
			Topic: topic,
		},
	}

	conn.Write(packet)
	conn.Read(call)

	return nil

}

func (p *Proxy) SubscribeForRangeQuery(topic, st, et string, call func(interface{})) error {

	conn := &infra.Connection{}
	err := conn.OpenTCP(p.addr)
	if err != nil {
		return err
	}

	packet := &Packet{
		Type: "sub",
		Config: &Config{
			Type: "range",
			Period: &Period{
				StartTime: st,
				EndTime:   et,
			},
		},
		Message: &Message{
			Topic: topic,
		},
	}

	conn.Write(packet)
	conn.ReaderWithTimeout(call)

	return nil

}
