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
		fmt.Println("recebeu range")
		fmt.Println(v)
	}

	proxy.SubscribeForRangeQuery(
		"topic1",
		"2018-06-11T00:00:00",
		"2018-06-12T12:00:00",
		call,
	)

}
