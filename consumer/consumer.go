package consumer

import (
	"log"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/goku321/pulsar-proto-x/schema/person"
	"github.com/syucream/avro-protobuf/pkg/serde"
)

func Consume(c pulsar.Client, topic, subscription, name, key string) {
	x, err := serde.NewSerDe(&person.Person{})
	if err != nil {
		log.Fatalf("failed to ser/deser: %s", err)
	}
	psConsumer := pulsar.NewProtoSchema(x.Codec.Schema(), nil)
	cons, err := c.Subscribe(pulsar.ConsumerOptions{
		Topic:            topic,
		SubscriptionName: subscription,
		Type:             pulsar.KeyShared,
		Name:             name,
		Schema:           psConsumer,
		KeySharedPolicy: &pulsar.KeySharedPolicy{
			Mode:                    pulsar.KeySharedPolicyModeAutoSplit,
			AllowOutOfOrderDelivery: false,
		},
		DLQ: &pulsar.DLQPolicy{
			MaxDeliveries:   1,
			DeadLetterTopic: "dlq",
		},
	})
	if err != nil {
		log.Fatalf("failed to subscribe to topic stream: %v", err)
	}

	defer cons.Close()

	for msg := range cons.Chan() {
		time.Sleep(time.Second * 0)
		log.Printf("message with key: %s", msg.Key())
		log.Printf("message value: %s", string(msg.Payload()))
		// if strings.Contains(msg.Key(), name) {
		msg.Nack(msg)
		log.Printf("Nacked message with key: %s", msg.Key())
		// }
		// time.Sleep(time.Second * 1)
	}
}
