package handlers

import (
	"MiddleApp/internal/domain"
	"MiddleApp/internal/repositories/interfaces"
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

type Handlers struct {
	Router *mux.Router
	Data   interfaces.Repository
}

func New(repo interfaces.Repository) *Handlers {
	router := mux.NewRouter()

	h := &Handlers{
		Router: router,
		Data:   repo,
	}

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			h.Read(w, r)
		case http.MethodPut:
			h.Update(w, r)
		case http.MethodPost:
			h.Create(w, r)
		case http.MethodDelete:
			h.Delete(w, r)
		default:
			log.Fatal("Unknown method")
		}
	})

	// TODO: Узнать про кажлый http метод

	return h
}

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
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Println("Unmarshal error")
		return
	}

	err, _ = h.Data.Create(user.Name, user.Surname, user.Age)
	if err != nil {
		log.Println("Data create function error")
		return
	}

	w.Write([]byte("User successfully added"))
}

func (h *Handlers) Read(w http.ResponseWriter, r *http.Request) {
	list := h.Data.Read()

	for _, value := range list {
		v, err := json.Marshal(value)
		if err != nil {
			log.Println("Read marshal error")

			return
		}

		w.Write(v)
	}
}

func (h *Handlers) Update(w http.ResponseWriter, r *http.Request) {
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
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Println("Unmarshal error")
		return
	}

	err, _ = h.Data.Update(user.Id, user.Name, user.Surname, user.Age)
	if err != nil {
		log.Println("Data update function error")
		return
	}

	w.Write([]byte("User successfully updated"))
}

type deleteStruct struct {
	Id uint32 `json:"id"`
}

func (h *Handlers) Delete(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if string(body) == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var delStr deleteStruct
	err = json.Unmarshal(body, &delStr)
	if err != nil {
		log.Println("Unmarshal error")
		return
	}

	err, _ = h.Data.Delete(delStr.Id)
	if err != nil {
		log.Println("Data deletes function error")
		return
	}
	log.Println(delStr.Id)

	w.Write([]byte("User was deleted"))
}
