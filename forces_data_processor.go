// Processes the Data Collection for each force.

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

func NewDefaultForcesDataProcessor() (collector *ForcesDataProcessor) {
	dataCollector := &HttpDataCollector{
		endpoint: forcesEndpoint,
	}
	collector = &ForcesDataProcessor{
		dataCollector: dataCollector,
	}
	return collector
}

// Transforms data byte array to a []Forces array
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
