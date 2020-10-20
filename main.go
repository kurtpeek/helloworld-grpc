package main

import (
	"fmt"
	"log"

	"github.com/kurtpeek/helloworld-grpc/helloworld"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	helloRequest := &helloworld.HelloRequest{Name: "John Doe"}
	helloRequestJSON, err := protojson.Marshal(helloRequest)
	if err != nil {
		log.Fatalf("Marshal: %v", err)
	}
	fmt.Println(string(helloRequestJSON))
}
