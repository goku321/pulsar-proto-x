package main

import (
	"flag"
	"log"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/goku321/pulsar-proto-x/producer"
)

func main() {
	sub := flag.String("sub", "test-sub", "subscription name")
	topic := flag.String("topic", "my-topic", "topic name")
	key := flag.String("key", "test-key", "key name")
	name := flag.String("name", "test-consumer", "consumer/producer name")
	mode := flag.String("mode", "consumer", "consumer/producer")
	count := flag.Int("count", 100, "message count")
	flag.Parse()

	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:               "pulsar://localhost:6650",
		OperationTimeout:  30 * time.Second,
		ConnectionTimeout: 30 * time.Second,
	})
	if err != nil {
		log.Fatalf("failed to init pulsar client: %s", err)
	}

	if *mode == "consumer" {
		// consume(client, *topic, *sub, *name, *key)
		log.Println(sub)
	} else {
		producer.Produce(client, *topic, *name, *key, *count)
	}
}
