package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/HistoireDeBabar/iso-string-converter"
	"sync"
	"time"
)

const form = "2006-Jan-02"

type UpdateChecker struct {
	policeDateCollector      DataCollector
	lastUpdatedDateCollector DataCollector
	policeUpdatedDate        time.Time
	lastCheckedDate          time.Time
	wg                       sync.WaitGroup
}

// Returns an UpdateChecker with the given DataCollectors.
func NewUpdateChecker(policeDataCollector DataCollector, lastUpdated DataCollector) (checker Checker) {
	checker = &UpdateChecker{
		policeDateCollector:      policeDataCollector,
		lastUpdatedDateCollector: lastUpdated,
	}
	return
}

// Returns a Default Update Checker with predefined configs.
func NewDefaultUpdateChecker() (checker Checker) {
	policeDataCollector := NewHttpDataCollector(dateCheckEndpoint)
	lastUpdated := NewDefaultS3DataCollector()
	checker = &UpdateChecker{
		policeDateCollector:      policeDataCollector,
		lastUpdatedDateCollector: lastUpdated,
	}
	return
}

// Internal Class for parsing data data.
type DateString struct {
	Date string
}

// Perform a Check to see whether the data needs to update.
func (dateChecker *UpdateChecker) Check() (valid bool, update time.Time) {
	dateChecker.wg.Add(2)
	go dateChecker.getPoliceData()
	go dateChecker.getLastUpdatedAt()
	dateChecker.wg.Wait()
	return dateChecker.CanUpdate(), dateChecker.policeUpdatedDate
}

func (dateChecker *UpdateChecker) getPoliceData() {
	policeUpdateData, _ := dateChecker.policeDateCollector.Collect()
	dateChecker.policeUpdatedDate, _ = dateChecker.TransformPoliceDate(policeUpdateData)
	fmt.Printf("Last Updated By Police: %v \n", dateChecker.policeUpdatedDate)
	dateChecker.wg.Done()
}

func (dateChecker *UpdateChecker) getLastUpdatedAt() {
	lastUpdatedAt, _ := dateChecker.lastUpdatedDateCollector.Collect()
	dateChecker.lastCheckedDate, _ = dateChecker.TransformLastUpdated(lastUpdatedAt)
	fmt.Printf("Last Updated By Crime Cross: %v \n", dateChecker.lastCheckedDate)
	dateChecker.wg.Done()
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
		fmt.Printf("%v \n", err.Error())
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
		fmt.Printf("%v \n", err.Error())
		return time.Time{}, err
	}
	if len(dates) == 0 {
		return time.Time{}, errors.New("No dates in last updated data array")
	}

	lastUpdatedAt, err = time.Parse(form, dates[0].Date)
	if err != nil {
		fmt.Printf("%v \n", err.Error())
		return time.Time{}, err
	}
	return lastUpdatedAt, nil
}
