package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	coap "github.com/go-ocf/go-coap"
)

// ArgsNum is the number of required arguments for application
const ArgsNum = 3

func main() {
	if len(os.Args) != ArgsNum {
		fmt.Printf("usage: coap-demo [address e.g. 192.168.73.192:1378] [path e.g. elahe]\n")
		return
	}

	co, err := coap.Dial("udp", os.Args[1])
	if err != nil {
		log.Fatalf("Error dialing: %v", err)
	}

	path := os.Args[2]

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	resp, err := co.GetWithContext(ctx, path)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}

	log.Printf("Response code: %s", resp.Code())

	var r map[string]interface{}
	if err := json.Unmarshal(resp.Payload(), &r); err != nil {
		log.Fatal(err)
	}

	log.Printf("Response payload: %+v", r)
}
