package bunny

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Bunny interface {
	Healthy() bool
}

type HttpBunny struct {
	url string
}

func NewBunny(url string) (bunny HttpBunny) {
	bunny.url = url
	return
}

func (bunny HttpBunny) Healthy() bool {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	res, err := client.Get(bunny.url)
	if err != nil {
		fmt.Println(err)
		return false
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		fmt.Println(err)
		return false
	}

	bunnyStatus := BunnyStatus{}
	json.Unmarshal(body, &bunnyStatus)
	if bunnyStatus.Status == "OK" {
		return true
	}
	return false
}

type BunnyStatus struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}
