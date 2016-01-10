package main

import (
	"encoding/json"
	"github.com/HistoireDeBabar/iso-string-converter"
	"time"
)

type UpdateChecker struct {
	dataCollector DataCollector
}

func NewUpdateChecker(dataCollector DataCollector) (collector *UpdateChecker) {
	collector = &UpdateChecker{
		dataCollector: dataCollector,
	}
	return collector
}

func NewDefaultUpdateChecker() (collector *UpdateChecker) {
	dataCollector := &HttpDataCollector{
		endpoint: dateCheckEndpoint,
	}
	collector = &UpdateChecker{
		dataCollector: dataCollector,
	}
	return collector
}

type DateString struct {
	Date string
}

func (dateCollector *UpdateChecker) CanUpdate(lastUpdatedAt time.Time, lastChecked time.Time) (valid bool) {
	if lastUpdatedAt.After(lastChecked) {
		return true
	}
	return false
}

func (dateCollector *UpdateChecker) Transform(data []byte) (updatedAt time.Time, err error) {
	dates := []DateString{}
	err = json.Unmarshal(data, &dates)
	if err != nil {
		return time.Time{}, err
	}
	return isoConverter.IsoStringToDate(dates[0].Date), nil
}
