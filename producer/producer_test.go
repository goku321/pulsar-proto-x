package producer

import (
	"context"
	"log"
	"testing"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/goku321/pulsar-proto-x/schema/person"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/syucream/avro-protobuf/pkg/serde"
)

func createClient() pulsar.Client {
	// create client
	lookupURL := "pulsar://localhost:6650"
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL: lookupURL,
	})
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func TestProducerWithSimpleSchema(t *testing.T) {
	c := createClient()
	p, err := serde.NewSerDe(&person.Person{})
	require.NoError(t, err)

	// Get the avro schema string.
	producerSchema := pulsar.NewProtoSchema(p.Codec.Schema(), nil)
	producer, err := c.CreateProducer(pulsar.ProducerOptions{
		Topic:  "proto_person",
		Schema: producerSchema,
	})
	require.Nil(t, err)

	if _, err := producer.Send(context.Background(), &pulsar.ProducerMessage{
		Value: &person.Person{
			Name: "luke",
			Age:  34,
		},
	}); err != nil {
		log.Fatal(err)
	}

	// Create consumer
	unobj := person.Person{}
	consumerSchema := pulsar.NewProtoSchema(p.Codec.Schema(), nil)
	consumer, err := c.Subscribe(pulsar.ConsumerOptions{
		Topic:                       "proto_person",
		SubscriptionName:            "sub-1",
		Schema:                      consumerSchema,
		SubscriptionInitialPosition: pulsar.SubscriptionPositionEarliest,
	})
	require.Nil(t, err)
	defer consumer.Close()

	msg, err := consumer.Receive(context.Background())
	assert.Nil(t, err)
	err = msg.GetSchemaValue(&unobj)
	require.Nil(t, err)
	assert.Equal(t, int32(34), unobj.Age)
	assert.Equal(t, "luke", unobj.Name)
}
