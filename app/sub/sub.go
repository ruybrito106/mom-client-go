package main

import (
	"fmt"
	"os"
	"time"

	dist "github.com/ruybrito106/mom-client-go/src/distribution"
)

func main() {

	proxy, err := dist.New("localhost:9091")
	if err != nil {
		os.Exit(1)
	}

	call := func(v interface{}) {
		fmt.Println("recebeu mizera")
		fmt.Println(v)
	}

	go func() {
		proxy.Subscribe("topic1", call)
	}()

	time.Sleep(100 * time.Second)

}
