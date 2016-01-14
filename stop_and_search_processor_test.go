package main

import (
	"github.com/HistoireDeBabar/crime-cross/mocks"
	"testing"
	"time"
)

func TestGenerateUrlFromTimeAndId(t *testing.T) {
	expectedUrl := "https://data.police.uk/api/stops-force?force=south_shields&date=2016-01"
	ti := time.Date(2016, time.January, 1, 1, 1, 1, 1, &time.Location{})
	id := "south_shields"
	subject := GenerateStopAndSearchUrl(ti, id)
	if subject != expectedUrl {
		t.Errorf("Error: subject generated: %v", subject)
	}
}

func TestGenerateUrlFromTimeAfterOctoberAndId(t *testing.T) {
	expectedUrl := "https://data.police.uk/api/stops-force?force=south_shields&date=2016-11"
	ti := time.Date(2016, time.November, 1, 1, 1, 1, 1, &time.Location{})
	id := "south_shields"
	subject := GenerateStopAndSearchUrl(ti, id)
	if subject != expectedUrl {
		t.Errorf("Error: subject generated: %v", subject)
	}
}

func TestTransformShouldReturnArrayOfStopAndSearched(t *testing.T) {
	mock := &mocks.MockStopAndSearchCollector{}
	subject := NewStopAndSearchProcessor(mock)
	response, _ := mock.Collect()
	results, _ := subject.Transform(response)
	if len(results) != 1 {
		t.Errorf("Expected length of results to equal 1, got :i %d", len(results))
	}
	result := results[0]
	if result.Legislation != "Misuse of Drugs Act 1971 (section 23)" {
		t.Errorf("Error In Name, got: %v", result.Legislation)
	}

	if result.DateTime != "2015-04-01T17:30:00" {
		t.Errorf("Error in DateTime, got: %v", result.DateTime)
	}
}
