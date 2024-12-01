package handlers

import (
	"Application/ValidationService/api/api/ServiceApiPb"
	"Application/ValidationService/internal/domain"
	"Application/ValidationService/utils"

	"context"
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

type Handlers struct {
	Router   *mux.Router
	Producer sarama.SyncProducer
	Consumer sarama.Consumer
	client   ServiceApiPb.UsersServiceClient
}

func New(producer sarama.SyncProducer, consumer sarama.Consumer, client ServiceApiPb.UsersServiceClient) *Handlers {
	h := &Handlers{}
	router := mux.NewRouter()
	h.Router = router
	h.Producer = producer
	h.Consumer = consumer
	h.client = client

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			h.Read(w, r)
		case http.MethodPut:
			h.Update(w, r)
		case http.MethodDelete:
			h.Delete(w, r)
		case http.MethodPost:
			h.Create(w, r)
		default:
			log.Fatal("unknown method")
		}
	})

	return h
}

// Create ------------------------
func (h *Handlers) Create(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if string(body) == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var user domain.User
	if err = json.Unmarshal(body, &user); err != nil {
		return
	}

	if err = utils.CreateValidation(user); err != nil {
		w.Write([]byte("Неверные данные" + err.Error()))
		return
	}

	msg := &sarama.ProducerMessage{
		Topic:     "CreateRequest",
		Partition: -1,
		Value:     sarama.ByteEncoder(body),
	}
	_, _, err = h.Producer.SendMessage(msg)
	if err != nil {
		error.Error(err)
	}

	claim, err := h.Consumer.ConsumePartition("CreateResponse", 0, sarama.OffsetOldest)
	if err != nil {
		error.Error(err)
	}

	select {
	case err = <-claim.Errors():
		log.Println(err)
	default:
	}

	w.Write([]byte("Юзер успешно добавлен"))
}

// Read ------------------------
type readRequest struct {
}

func (h *Handlers) Read(w http.ResponseWriter, r *http.Request) {
	var in *ServiceApiPb.ReadRequest
	ctx := context.Background()
	resp, err := h.client.Read(ctx, in)
	if err != nil {
		fmt.Printf(err.Error())
	}

	for _, student := range resp.Users {
		val, _ := json.Marshal(student)
		w.Write(val)
	}
}

// Update ------------------------
func (h *Handlers) Update(w http.ResponseWriter, r *http.Request) {

}

// Delete ------------------------
type deleteRequest struct {
	Id uint `json:"id"`
}

func (h *Handlers) Delete(w http.ResponseWriter, r *http.Request) {

}
