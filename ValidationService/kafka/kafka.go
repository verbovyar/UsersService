package kafka

import (
	"github.com/Shopify/sarama"
	"log"
	"os"
)

var brokers = []string{"127.0.0.1:9092"}

func NewProducer() sarama.SyncProducer {
	conf := sarama.NewConfig()
	conf.Version = sarama.V1_1_0_0
	conf.Producer.Partitioner = sarama.NewRandomPartitioner
	conf.Producer.RequiredAcks = sarama.WaitForLocal
	conf.Producer.Return.Successes = true

	sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)

	producer, err := sarama.NewSyncProducer(brokers, conf)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return producer
}

func NewConsumer() sarama.Consumer {
	conf := sarama.NewConfig()
	//conf.Version = sarama.V0_10_0_0

	sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)
	consumer, err := sarama.NewConsumer(brokers, conf)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return consumer
}
