package routes

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"testing"

	"github.com/ha-dev/getslowestdbqueries/env"
	"github.com/stretchr/testify/assert"
	//"github.com/stretchr/testify/assert"
)

type ResultBody struct {
	Error             bool        `json:"error"`
	Msg               string      `json:"msg"`
	Count             uint64      `json:"count"`
	ListSlowestQuerys interface{} `json:"listSlowestQuerys"`
}

func TestGetQuerys(t *testing.T) {

	endpoint := "http://" + env.ServicePort + "/api/v1/List"
	data := url.Values{}
	data.Set("query_command", "select")
	data.Set("rowcountpage", "2")
	data.Set("currentpage", "1")

	client := &http.Client{}
	r, err := http.NewRequest("POST", endpoint, strings.NewReader(data.Encode())) // URL-encoded payload
	if err != nil {
		t.Log(err.Error())
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, err := client.Do(r)
	if err != nil {
		t.Log(err.Error())
	}

	defer resp.Body.Close()
	v := &ResultBody{}
	err = json.NewDecoder(resp.Body).Decode(v) //decode the request body into struct and failed if any error occur
	//t.Log(resp)
	if err != nil {
		t.Log(err.Error())
	}
	assert.Equal(t, false, v.Error, "they should be equal")

}

func TestAddQuerysForCheckSlowest(t *testing.T) {

	endpoint := "http://" + env.ServicePort + "/api/v1/AddQueryForCheck"
	data := url.Values{}
	query := `INSERT INTO users (username, password)
	    VALUES('test','1');`
	data.Set("query", query)
	client := &http.Client{}
	r, err := http.NewRequest("POST", endpoint, strings.NewReader(data.Encode())) // URL-encoded payload
	if err != nil {
		t.Log(err.Error())
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, err := client.Do(r)
	if err != nil {
		t.Log(err.Error())
	}

	defer resp.Body.Close()
	v := &ResultBody{}
	err = json.NewDecoder(resp.Body).Decode(v) //decode the request body into struct and failed if any error occur
	//t.Log(resp)
	if err != nil {
		t.Log(err.Error())
	}
	assert.Equal(t, false, v.Error, "they should be equal")

}
