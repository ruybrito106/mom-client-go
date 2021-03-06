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

	for i := 10; i < 20; i++ {
		err = proxy.Publish("topic1", fmt.Sprintf("from_golang_%d", i))
		if err == nil {
			fmt.Println("Msg published successfully")
		}
	}

	time.Sleep(1 * time.Second)

}
