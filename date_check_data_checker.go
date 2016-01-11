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
	policeUpdatedDate        time.Time
	lastCheckedDate          time.Time
}

// Returns an UpdateChecker with the given DataCollectors.
func NewUpdateChecker(policeDataCollector DataCollector, lastUpdated DataCollector) (collector *UpdateChecker) {
	collector = &UpdateChecker{
		policeDateCollector:      policeDataCollector,
		lastUpdatedDateCollector: lastUpdated,
	}
	return collector
}

// Returns a Default Update Checker with predefined configs.
func NewDefaultUpdateChecker() (collector Checker) {
	policeDataCollector := NewHttpDataCollector(dateCheckEndpoint)
	lastUpdated := NewDefaultS3DataCollector()
	collector = &UpdateChecker{
		policeDateCollector:      policeDataCollector,
		lastUpdatedDateCollector: lastUpdated,
	}
	return collector
}

// Internal Class for parsing data data.
type DateString struct {
	Date string
}

// Perform a Check to see whether the data needs to update.
func (dateChecker *UpdateChecker) Check() (valid bool) {
	go dateChecker.getPoliceData()
	go dateChecker.getLastUpdatedAt()
	return dateChecker.CanUpdate()
}

func (dateChecker *UpdateChecker) getPoliceData() {
	policeUpdateData, _ := dateChecker.policeDateCollector.Collect()
	dateChecker.policeUpdatedDate, _ = dateChecker.TransformPoliceDate(policeUpdateData)
}

func (dateChecker *UpdateChecker) getLastUpdatedAt() {
	lastUpdatedAt, _ := dateChecker.lastUpdatedDateCollector.Collect()
	dateChecker.lastCheckedDate, _ = dateChecker.TransformLastUpdated(lastUpdatedAt)
}

// Date compare function.
func (dateChecker *UpdateChecker) CanUpdate() (valid bool) {
	//fallback where if either is zero we can update
	if dateChecker.policeUpdatedDate.IsZero() || dateChecker.lastCheckedDate.IsZero() {
		return true
	}

	if dateChecker.policeUpdatedDate.After(dateChecker.lastCheckedDate) {
		return true
	}
	return false
}

// Transform Method Transforms a ByteArray in the format of the PoliceApi
// and returns a Time object.
func (dateChecker *UpdateChecker) TransformPoliceDate(data []byte) (updatedAt time.Time, err error) {
	dates := DateString{}
	err = json.Unmarshal(data, &dates)
	if err != nil {
		fmt.Println(err)
		return time.Time{}, err
	}
	return isoConverter.IsoStringToDate(dates.Date), nil
}

// Transform Method Transforms a ByteArray in the format of the S3
// and returns a Time object.
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
