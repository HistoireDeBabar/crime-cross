package main

import (
	"io/ioutil"
	"net/http"
)

type HttpDataCollector struct {
	client   *http.Client
	endpoint string
}

func NewHttpDataCollector(endpoint string) (collector DataCollector) {
	httpDataCollector := &HttpDataCollector{
		client:   &http.Client{},
		endpoint: endpoint,
	}
	return httpDataCollector
}

func (dataCollector *HttpDataCollector) Collect() (data []byte, err error) {
	response, err := dataCollector.client.Get(dataCollector.endpoint)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return data, nil
	}
	return body, nil
}
