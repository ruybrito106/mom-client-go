package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	dist "github.com/ruybrito106/mom-client-go/src/distribution"
)

func main() {

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	proxy, err := dist.New("localhost:9091")
	if err != nil {
		os.Exit(1)
	}

	call := func(v interface{}) {
		fmt.Println("recebeu real time")
		fmt.Println(v)
	}

	go func() {
		proxy.Subscribe(
			"topic1",
			call,
		)
	}()

	<-signals
	os.Exit(0)

}
