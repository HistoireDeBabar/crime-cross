package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	forceQS = "force="
	and     = "&"
	timeQS  = "date="
	join    = "-"
)

type StopAndSearchProcessor struct {
	dataCollector DataCollector
}

func NewStopAndSearchProcessor(dataCollector DataCollector) *StopAndSearchProcessor {
	return &StopAndSearchProcessor{
		dataCollector: dataCollector,
	}
}

func NewStopAndSearchProcessorFromTimeAndForceId(time time.Time, forceId string) *StopAndSearchProcessor {
	url := GenerateStopAndSearchUrl(time, forceId)
	dataCollector := NewHttpDataCollector(url)
	return &StopAndSearchProcessor{
		dataCollector: dataCollector,
	}
}

func (s *StopAndSearchProcessor) GetStopAndSearches() (stopAndSearches []StopAndSearch) {
	data, err := s.dataCollector.Collect()
	if err != nil {
		return stopAndSearches
	}
	stopAndSearches, err = s.Transform(data)
	return
}

func (s *StopAndSearchProcessor) Transform(data []byte) (stopAndSearches []StopAndSearch, err error) {
	if len(data) == 0 {
		return stopAndSearches, nil
	}

	err = json.Unmarshal(data, &stopAndSearches)
	if err != nil {
		fmt.Printf("%v", err.Error())
		return nil, err
	}
	return stopAndSearches, nil
}

func GenerateStopAndSearchUrl(time time.Time, id string) (url string) {
	year := int(time.Year())
	month := formatMonth(int(time.Month()))
	path := []string{
		stopAndSearchEndpoint,
		forceQS,
		id,
		and,
		timeQS,
		strconv.Itoa(year),
		join,
		month,
	}
	url = strings.Join(path, "")
	return url
}

func formatMonth(month int) (rep string) {
	if month < 10 {
		return "0" + strconv.Itoa(month)
	}
	return strconv.Itoa(month)
}
