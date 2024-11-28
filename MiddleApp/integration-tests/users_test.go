package integration_tests

import (
	"MiddleApp/config"
	"MiddleApp/internal/app"
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"net/http"
	"testing"
	"time"
)

func TestDelete(t *testing.T) {
	conf, err := config.LoadConfig("./config")
	if err != nil {
		println(err.Error())
	}

	repo := app.RunRepo(&conf)
	go app.RunHttp(&conf, repo)

	client := http.Client{Timeout: 6 * time.Second}

	body := struct {
		id uint32
	}{id: 2}
	reqBody, err := json.Marshal(body)
	req, err := http.NewRequest("DELETE", "http://localhost:9000", bytes.NewBuffer(reqBody))
	if err != nil {
		log.Println(err.Error())
	}

	resp, err := client.Do(req)
	respBody, err := io.ReadAll(resp.Body)
	user := struct {
		id uint32
	}{}
	err = json.Unmarshal(respBody, &user)
	if err == nil && user.id == body.id {
		assert.True(t, true)
	} else {
		assert.True(t, false)
	}

}
