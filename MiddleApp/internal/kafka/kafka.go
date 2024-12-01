package kafka

import (
	"github.com/IBM/sarama"
	"log"
	"os"
)

var brokers = []string{"127.0.0.1:9092"}

func NewProducer() sarama.SyncProducer {
	conf := sarama.NewConfig()
	conf.Version = sarama.V0_11_0_0
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

func NewConsumerGroup() sarama.ConsumerGroup {
	conf := sarama.NewConfig()
	conf.Version = sarama.V0_11_0_0
	conf.Consumer.Return.Errors = true
	conf.Consumer.Offsets.Initial = sarama.OffsetOldest

	sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)

	consumerGroup, err := sarama.NewConsumerGroup(brokers, "Start consuming on 9092 port", conf)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return consumerGroup
}
