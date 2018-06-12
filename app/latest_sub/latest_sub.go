package main

import (
	"fmt"
	"os"

	dist "github.com/ruybrito106/mom-client-go/src/distribution"
)

func main() {

	proxy, err := dist.New("localhost:9091")
	if err != nil {
		os.Exit(1)
	}

	call := func(v interface{}) {
		fmt.Println("recebeu latest")
		fmt.Println(v)
	}

	proxy.SubscribeForLatest(
		"topic1",
		call,
	)

}
