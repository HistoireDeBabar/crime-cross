package main

import (
	"github.com/HistoireDeBabar/crime-cross/mocks"
	"net/http"
	"testing"
)

func TestResponseWithNoBodyReturnsEmptyArray(t *testing.T) {
	mockClient := http.Client{
		Transport: &mocks.SuccessForceListNoResponseBody{
			StatusCode: 200,
		},
	}

	subject := &HttpDataCollector{
		client:   &mockClient,
		endpoint: "https://data.police.uk/api/forces",
	}
	data, err := subject.Collect()
	if err != nil {
		t.Errorf("Unexpected Error %v", err)
	}
	if len(data) != 0 {
		t.Errorf("expected result length to be 0 got: %v", len(data))
	}
}

func TestReturnsValidByteArrary(t *testing.T) {
	mockClient := http.Client{
		Transport: &mocks.SuccessForceListWithResponseBody{
			StatusCode: 200,
		},
	}

	dataCollector := &HttpDataCollector{
		client:   &mockClient,
		endpoint: "https://data.police.uk/api/forces",
	}

	result, _ := dataCollector.Collect()

	expectedResult := []byte{91, 123, 34, 105, 100, 34, 58, 32, 34, 115, 111, 117, 116, 104, 95, 115, 104, 105, 101, 108, 100, 115, 34, 44, 32, 34, 110, 97, 109, 101, 34, 58, 32, 34, 83, 111, 117, 116, 104, 32, 83, 104, 105, 101, 108, 100, 115, 34, 125, 93}
	if result == nil {
		t.Errorf("Expected result to equal %v got %v", result, expectedResult)
	}
	if len(result) != len(expectedResult) {
		t.Errorf("Unexpected Error: Response is incorrect lenght")
	}
}
