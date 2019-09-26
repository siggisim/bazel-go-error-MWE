package main

import (
	"fmt"
	"cloud.google.com/go/pubsub"
)

func main() {
	c := &pubsub.Client{}
	fmt.Printf("hello %p", c)
}


