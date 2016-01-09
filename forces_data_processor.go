// Retrieves data from the police api about specific forces

package main

import (
	"encoding/json"
)

type ForcesDataProcessor struct {
	dataCollector DataCollector
}

func NewForcesDataProcessor(dataCollector DataCollector) (transformer *ForcesDataProcessor) {
	transformer = &ForcesDataProcessor{
		dataCollector: dataCollector,
	}
	return transformer
}

// Retreieves all the basic data about the forces from the police api
func (fdc *ForcesDataProcessor) Transform(data []byte) (forces []Force, err error) {
	if len(data) == 0 {
		return forces, nil
	}

	err = json.Unmarshal(data, &forces)
	if err != nil {
		return nil, err
	}

	return forces, nil
}
