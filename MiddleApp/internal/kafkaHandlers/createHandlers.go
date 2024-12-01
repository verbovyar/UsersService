package kafkaHandlers

import (
	"MiddleApp/internal/repositories/interfaces"
	"context"
	"encoding/json"
	"github.com/IBM/sarama"
	"log"
	"time"
)

type CreateHandler struct {
	producer      sarama.SyncProducer
	consumerGroup sarama.ConsumerGroup
	data          interfaces.Repository
}

func NewCreateHandler(producer sarama.SyncProducer, consumerGroup sarama.ConsumerGroup, data interfaces.Repository) *CreateHandler {
	return &CreateHandler{
		producer:      producer,
		consumerGroup: consumerGroup,
		data:          data,
	}
}

type addRequest struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     uint32 `json:"age"`
}

type addResponse struct {
	Error error  `json:"error"`
	Id    uint32 `json:"id"`
}

func (h *CreateHandler) Setup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (h *CreateHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (h *CreateHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		var request addRequest
		err := json.Unmarshal(msg.Value, &request)
		if err != nil {
			log.Printf("income data %v: %v", string(msg.Value), err)
			continue
		}

		err, id := h.data.Create(request.Name, request.Surname, request.Age)

		addResp := addResponse{
			Id:    id,
			Error: err,
		}
		response, _ := json.Marshal(&addResp)

		producerMsg := &sarama.ProducerMessage{
			Topic:     "CreateResponse",
			Partition: -1,
			Value:     sarama.ByteEncoder(response),
		}

		_, _, err = h.producer.SendMessage(producerMsg)

	}

	return nil
}

func CreateClaim(createHandler *CreateHandler) {
	for {
		if err := createHandler.consumerGroup.Consume(context.Background(), []string{"CreateRequest"}, createHandler); err != nil {
			log.Printf("on consume: %v", err)
			time.Sleep(time.Second * 10)
		}
	}
}
