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

type BunnyStatus struct {
	Name   string `json:"name"`
	Status string `json:"status"`
	ProcessTime int `json:"timeToEvaluate"`
}

func NewBunny(url string) (bunny HttpBunny) {
	bunny.url = url
	return
}


func(bunny HttpBunny) Update() (bunnyStatus BunnyStatus) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}


	res, err := client.Get(bunny.url)
	if err != nil {
		fmt.Println(err)
		bunnyStatus.Status = "ERROR"
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		fmt.Println(err)
		bunnyStatus.Status = "ERROR"
	}else{
		json.Unmarshal(body, &bunnyStatus)
	}
	//fmt.Println(string(body))
	
	return
}

