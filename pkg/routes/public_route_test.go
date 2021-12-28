package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"testing"

	"github.com/ha-dev/getslowestdbqueries/env"
	"github.com/ha-dev/getslowestdbqueries/models"
	"github.com/stretchr/testify/assert"
)

type ResultBody struct {
	Error             bool                 `json:"error"`
	Msg               string               `json:"msg"`
	Count             uint64               `json:"count"`
	ListSlowestQuerys []models.ResulyQuery `json:"listSlowestQuerys"`
}

func TestGetQuerys(t *testing.T) {
	var client http.Client
	resp, err := client.Get("http://" + env.ServicePort + "/api/v1/list")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	v := &ResultBody{}
	err = json.NewDecoder(resp.Body).Decode(v) //decode the request body into struct and failed if any error occur
	if err != nil {
		t.Log(err.Error())
	}
	assert.Equal(t, false, v.Error, "they should be equal")
}
