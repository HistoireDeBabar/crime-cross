package forces

import (
	"github.com/HistoireDeBabar/crime-cross/mocks"
	"net/http"
	"testing"
)

func TestRequestReturnsNoResponseInBodReturnsNoForces(t *testing.T) {
	mockClient := http.Client{
		Transport: &mocks.SuccessForceListNoResponseBody{
			StatusCode: 200,
		},
	}

	subject := ForcesDataCollector{
		Client:   &mockClient,
		endpoint: "https://data.police.uk/api/forces",
	}

	result, err := subject.GetForcesIdentifier()

	if err != nil {
		t.Errorf("Unexpected Error %v", err)
	}
	if len(result) != 0 {
		t.Errorf("expected result length to be 0 got: %v", len(result))
	}
}

func TestRequestReturnsValidBody(t *testing.T) {
	mockClient := http.Client{
		Transport: &mocks.SuccessForceListWithResponseBody{
			StatusCode: 200,
		},
	}

	subject := ForcesDataCollector{
		Client:   &mockClient,
		endpoint: "https://data.police.uk/api/forces",
	}

	result, err := subject.GetForcesIdentifier()

	if err != nil {
		t.Errorf("Unexpected Error %v", err)
	}
	if len(result) != 1 {
		t.Errorf("expected result length to be 1 got: %v", len(result))
	}
	var force = result[0]
	if force.Id != "south_shields" {
		t.Errorf("expected id to equal south_shields got", force.Id)
	}

	if force.Name != "South Shields" {
		t.Errorf("expected id to equal South Shields got", force.Name)
	}
}

func TestClientErrorCode(t *testing.T) {
	mockClient := http.Client{
		Transport: &mocks.ResponseErrorStatusCode{
			StatusCode: 500,
		},
	}

	subject := ForcesDataCollector{
		Client:   &mockClient,
		endpoint: "https://data.police.uk/api/forces",
	}

	_, err := subject.GetForcesIdentifier()
	if err == nil {
		t.Errorf("Expected Error from Server, got Nil")
	}
	if err.Error() != "Error Response from Server" {
		t.Errorf("Unexpected ErrorResponse %v", err)
	}
}

func TestErrorFromServerServerError(t *testing.T) {
	mockClient := http.Client{
		Transport: &mocks.ResponseErrorStatusCode{
			StatusCode: 404,
		},
	}

	subject := ForcesDataCollector{
		Client:   &mockClient,
		endpoint: "https://data.police.uk/api/forces",
	}

	_, err := subject.GetForcesIdentifier()
	if err == nil {
		t.Errorf("Expected Error from Server, got Nil")
	}
	if err.Error() != "Error Response from Server" {
		t.Errorf("Unexpected ErrorResponse %v", err)
	}
}
