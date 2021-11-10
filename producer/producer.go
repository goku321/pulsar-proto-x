package producer

import (
	"context"
	"log"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/goku321/pulsar-proto-x/schema/person"
	"github.com/syucream/avro-protobuf/pkg/serde"
)

func noop(id pulsar.MessageID, msg *pulsar.ProducerMessage, err error) {
	if err != nil {
		log.Println("message produced a.h.a!")
	}
}

func Produce(c pulsar.Client, topic, name, key string, count int) {
	x, err := serde.NewSerDe(&person.Person{})
	if err != nil {
		log.Fatalf("failed to ser/deser: %s", err)
	}
	psProducer := pulsar.NewProtoSchema(x.Codec.Schema(), nil)
	p, err := c.CreateProducer(pulsar.ProducerOptions{
		Topic:           topic,
		Name:            name,
		DisableBatching: true,
		Schema:          psProducer,
	})
	if err != nil {
		log.Fatalf("failed to create producer: %s", err)
	}

	for i := 0; i < count; i++ {
		// p.SendAsync(context.Background(), &pulsar.ProducerMessage{
		// 	Payload: []byte("hello pulsar"),
		// 	Key: key,
		// }, noop)

		if _, err := p.Send(context.Background(), &pulsar.ProducerMessage{
			Value: &person.Person{
				Name: "Luke Skywalker",
				Age:  34,
			},
			Key: key,
		}); err != nil {
			log.Fatal("failed to produce message")
		}
	}
}
