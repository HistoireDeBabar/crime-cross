package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/HistoireDeBabar/iso-string-converter"
	"time"
)

const form = "2006-Jan-02"

type UpdateChecker struct {
	policeDateCollector      DataCollector
	lastUpdatedDateCollector DataCollector
}

func NewUpdateChecker(policeDataCollector DataCollector, lastUpdated DataCollector) (collector *UpdateChecker) {
	collector = &UpdateChecker{
		policeDateCollector:      policeDataCollector,
		lastUpdatedDateCollector: lastUpdated,
	}
	return collector
}

func NewDefaultUpdateChecker() (collector Checker) {
	policeDataCollector := NewHttpDataCollector(dateCheckEndpoint)
	lastUpdated := NewDefaultS3DataCollector()
	collector = &UpdateChecker{
		policeDateCollector:      policeDataCollector,
		lastUpdatedDateCollector: lastUpdated,
	}
	return collector
}

type DateString struct {
	Date string
}

func (dateChecker *UpdateChecker) Check() (valid bool) {
	//do data collection async
	policeUpdateData, err := dateChecker.policeDateCollector.Collect()
	if err != nil {
		fmt.Println(err)
		return false
	}
	lastUpdated, err := dateChecker.lastUpdatedDateCollector.Collect()
	if err != nil {
		fmt.Println(err)
		return false
	}
	policeUpdatedAtDate, _ := dateChecker.TransformPoliceDate(policeUpdateData)
	lastCheckedAtDate, _ := dateChecker.TransformLastUpdated(lastUpdated)
	return dateChecker.CanUpdate(policeUpdatedAtDate, lastCheckedAtDate)
}

func (dateChecker *UpdateChecker) CanUpdate(lastUpdatedAt time.Time, lastChecked time.Time) (valid bool) {
	//fallback where if either is zero we can update
	if lastUpdatedAt.IsZero() || lastChecked.IsZero() {
		return true
	}

	if lastUpdatedAt.After(lastChecked) {
		return true
	}
	return false
}

func (dateChecker *UpdateChecker) TransformPoliceDate(data []byte) (updatedAt time.Time, err error) {
	dates := DateString{}
	err = json.Unmarshal(data, &dates)
	if err != nil {
		fmt.Println(err)
		return time.Time{}, err
	}
	return isoConverter.IsoStringToDate(dates.Date), nil
}

func (dateChecker *UpdateChecker) TransformLastUpdated(data []byte) (lastUpdatedAt time.Time, err error) {
	dates := []DateString{}
	err = json.Unmarshal(data, &dates)
	if err != nil {
		return time.Time{}, err
	}
	if len(dates) == 0 {
		return time.Time{}, errors.New("No Date Data From S3")
	}
	lastUpdatedAt, err = time.Parse(form, dates[0].Date)
	if err != nil {
		return time.Time{}, err
	}
	return lastUpdatedAt, nil
}
