package producer

import (
	"context"
	"log"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/goku321/pulsar-proto-x/schema/example"
	"github.com/syucream/avro-protobuf/pkg/serde"
)

func noop(id pulsar.MessageID, msg *pulsar.ProducerMessage, err error) {
	if err != nil {
		log.Println("message produced a.h.a!")
	}
}

func Produce(c pulsar.Client, topic, name, key string, count int) {
	x, err := serde.NewSerDe(&example.Example{})
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
			Value: &example.Example{
				Uuid: "1",
				SingleNested: &example.Example_Nested{
					Name: "Luke",
					Age:  34,
					Ok:   example.Example_Nested_TRUE,
				},
			},
			Key: key,
		}); err != nil {
			log.Fatal("failed to produce message")
		}
	}
}
