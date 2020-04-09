package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	coap "github.com/go-ocf/go-coap"
)

const argsNum = 3

func main() {
	if len(os.Args) != argsNum {
		fmt.Printf("usage: coap-demo [address] [path]\n")
		return
	}

	co, err := coap.Dial("udp", os.Args[1])
	if err != nil {
		log.Fatalf("Error dialing: %v", err)
	}

	path := os.Args[2]

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := co.GetWithContext(ctx, path)

	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}

	log.Printf("Response payload: %s", string(resp.Payload()))
}
