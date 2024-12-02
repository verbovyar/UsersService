package kafkaHandlers

import (
	"MiddleApp/internal/repositories/interfaces"
	"context"
	"encoding/json"
	"github.com/IBM/sarama"
	"log"
	"time"
)

type UpdateHandler struct {
	producer      sarama.SyncProducer
	consumerGroup sarama.ConsumerGroup
	data          interfaces.Repository
}

func NewUpdateHandler(producer sarama.SyncProducer, consumerGroup sarama.ConsumerGroup, data interfaces.Repository) *UpdateHandler {
	return &UpdateHandler{
		producer:      producer,
		consumerGroup: consumerGroup,
		data:          data,
	}
}

type updateRequest struct {
	Id      uint32 `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     uint32 `json:"age"`
}

type updateResponse struct {
	Error error  `json:"error"`
	Id    uint32 `json:"id"`
}

func (h *UpdateHandler) Setup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (h *UpdateHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (h *UpdateHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		var request updateRequest
		err := json.Unmarshal(msg.Value, &request)
		if err != nil {
			log.Printf("income data %v: %v", string(msg.Value), err)
			continue
		}

		err, id := h.data.Update(request.Id, request.Name, request.Surname, request.Age)

		updateResp := updateResponse{
			Id:    id,
			Error: err,
		}
		response, _ := json.Marshal(&updateResp)

		producerMsg := &sarama.ProducerMessage{
			Topic:     "UpdateResponse",
			Partition: -1,
			Value:     sarama.ByteEncoder(response),
		}

		_, _, err = h.producer.SendMessage(producerMsg)

	}

	return nil
}

func UpdateClaim(updateHandler *UpdateHandler) {
	for {
		if err := updateHandler.consumerGroup.Consume(context.Background(), []string{"UpdateRequest"}, updateHandler); err != nil {
			log.Printf("on consume: %v", err)
			time.Sleep(time.Second * 10)
		}
	}
}
