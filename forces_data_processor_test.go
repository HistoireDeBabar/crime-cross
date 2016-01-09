package main

import (
	"testing"
)

func TestRequestReturnsNoResponseInBodReturnsNoForces(t *testing.T) {
	subject := ForcesDataProcessor{}
	result, err := subject.Transform([]byte{})

	if err != nil {
		t.Errorf("Unexpected Error %v", err)
	}
	if len(result) != 0 {
		t.Errorf("expected result length to be 0 got: %v", len(result))
	}
}

func TestRequestReturnsValidBody(t *testing.T) {

	subject := ForcesDataProcessor{}
	validResponse := []byte{91, 123, 34, 105, 100, 34, 58, 32, 34, 115, 111, 117, 116, 104, 95, 115, 104, 105, 101, 108, 100, 115, 34, 44, 32, 34, 110, 97, 109, 101, 34, 58, 32, 34, 83, 111, 117, 116, 104, 32, 83, 104, 105, 101, 108, 100, 115, 34, 125, 93}
	result, err := subject.Transform(validResponse)

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
