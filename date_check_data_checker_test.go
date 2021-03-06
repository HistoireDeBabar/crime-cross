package main

import (
	"github.com/HistoireDeBabar/crime-cross/mocks"
	"net/http"
	"testing"
	"time"
)

func TestGetValidDataFromRequest(t *testing.T) {
	mockClient := http.Client{
		Transport: &mocks.ValidDateTransport{
			StatusCode: 200,
		},
	}

	dataCollector := &HttpDataCollector{
		client:   &mockClient,
		endpoint: "https://data.police.uk/api/crime-last-updated",
	}
	data, _ := dataCollector.Collect()
	subject := UpdateChecker{}
	result, err := subject.TransformPoliceDate(data)
	if err != nil {
		t.Errorf("UnexpectedError:", err)
	}
	if result.IsZero() {
		t.Error("Result should be not be Zero time")
	}
	if result.Month() != time.January {
		t.Errorf("result Month should be January not %v", result.Month())
	}

	if result.Day() != 1 {
		t.Errorf("result Day should be 1 not %v", result.Day())
	}

	if result.Year() != 2011 {
		t.Errorf("result Day should be 2011 not %v", result.Year())
	}
}

func TestGetInValidDataFromRequest(t *testing.T) {
	mockClient := http.Client{
		Transport: &mocks.InValidDateTransport{
			StatusCode: 200,
		},
	}

	dataCollector := &HttpDataCollector{
		client:   &mockClient,
		endpoint: "https://data.police.uk/api/crime-last-updated",
	}
	data, _ := dataCollector.Collect()
	subject := UpdateChecker{}
	result, err := subject.TransformPoliceDate(data)
	if err != nil {
		t.Errorf("UnexpectedError:", err)
	}
	if !result.IsZero() {
		t.Error("Result should be not be Zero time")
	}
}

func TestWhetherADateCanBeUpdatedTruthy(t *testing.T) {
	lastUpdated := time.Date(2009, time.November, 10, 29, 0, 0, 0, time.UTC)
	lastChecked := time.Date(2009, time.November, 10, 24, 0, 0, 0, time.UTC)
	subject := UpdateChecker{
		policeUpdatedDate: lastUpdated,
		lastCheckedDate:   lastChecked,
	}
	canUpdate := subject.CanUpdate()
	if canUpdate == false {
		t.Errorf("Expected to be allowed to update")
	}
}

func TestWhetherADateCanBeUpdatedFalsey(t *testing.T) {

	lastUpdated := time.Date(2009, time.January, 10, 29, 0, 0, 0, time.UTC)
	lastChecked := time.Date(2009, time.November, 10, 24, 0, 0, 0, time.UTC)
	subject := UpdateChecker{
		policeUpdatedDate: lastUpdated,
		lastCheckedDate:   lastChecked,
	}
	canUpdate := subject.CanUpdate()
	if canUpdate == true {
		t.Errorf("Expected to not be allowed to update")
	}
}

func TestWhenEitherIsZeroReturnsTrue(t *testing.T) {
	lastUpdated := time.Time{}
	lastChecked := time.Date(2009, time.November, 10, 24, 0, 0, 0, time.UTC)
	subject := UpdateChecker{
		policeUpdatedDate: lastUpdated,
		lastCheckedDate:   lastChecked,
	}
	canUpdate := subject.CanUpdate()
	if canUpdate == false {
		t.Errorf("Expected to be true when last updated is zero")
	}

	lastUpdated = time.Date(2009, time.November, 10, 24, 0, 0, 0, time.UTC)
	lastChecked = time.Time{}

	subject = UpdateChecker{
		policeUpdatedDate: lastUpdated,
		lastCheckedDate:   lastChecked,
	}

	canUpdate = subject.CanUpdate()
	if canUpdate == false {
		t.Errorf("Expected to be true when last checked is zero")
	}
}

func TestTransformFromS3Document(t *testing.T) {
	data := []byte{
		91, 10, 32, 32, 123, 10, 32, 32, 32, 32, 34, 100,
		97, 116, 101, 34, 58, 32, 34, 50, 48, 49, 53, 45,
		77, 97, 114, 45, 50, 57, 34, 10, 32, 32, 125, 10, 93,
		10}
	subject := UpdateChecker{}
	result, err := subject.TransformLastUpdated(data)
	if err != nil {
		t.Errorf("UnexpectedError: %v", err)
	}

	if result.Year() != 2015 {
		t.Errorf("Year Expected 2015, Got %d", result.Year())
	}

	if result.IsZero() {
		t.Error("Result should be not be Zero time")
	}
	if result.Month() != time.March {
		t.Errorf("result Month should be March not %v", result.Month())
	}

	if result.Day() != 29 {
		t.Errorf("result Day should be 29 not %d", result.Day())
	}
}
