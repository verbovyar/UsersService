package kafkaHandlers

import (
	"MiddleApp/internal/repositories/interfaces"
	"context"
	"encoding/json"
	"github.com/IBM/sarama"
	"log"
	"time"
)

type DeleteHandler struct {
	producer      sarama.SyncProducer
	consumerGroup sarama.ConsumerGroup
	data          interfaces.Repository
}

func NewDeleteHandler(producer sarama.SyncProducer, consumerGroup sarama.ConsumerGroup, data interfaces.Repository) *DeleteHandler {
	return &DeleteHandler{
		producer:      producer,
		consumerGroup: consumerGroup,
		data:          data,
	}
}

type deleteRequest struct {
	Id uint32 `json:"id"`
}

type deleteResponse struct {
	Error error  `json:"error"`
	Id    uint32 `json:"id"`
}

func (h *DeleteHandler) Setup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (h *DeleteHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (h *DeleteHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		var request deleteRequest
		err := json.Unmarshal(msg.Value, &request)
		if err != nil {
			log.Printf("income data %v: %v", string(msg.Value), err)
			continue
		}

		err, id := h.data.Delete(request.Id)

		deleteResp := deleteResponse{
			Id:    id,
			Error: err,
		}
		response, _ := json.Marshal(&deleteResp)

		producerMsg := &sarama.ProducerMessage{
			Topic:     "DeleteResponse",
			Partition: -1,
			Value:     sarama.ByteEncoder(response),
		}

		_, _, err = h.producer.SendMessage(producerMsg)

	}

	return nil
}

func DeleteClaim(deleteHandler *DeleteHandler) {
	for {
		if err := deleteHandler.consumerGroup.Consume(context.Background(), []string{"DeleteRequest"}, deleteHandler); err != nil {
			log.Printf("on consume: %v", err)
			time.Sleep(time.Second * 10)
		}
	}
}
