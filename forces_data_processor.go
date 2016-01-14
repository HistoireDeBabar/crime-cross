// Processes the Data Collection for each force.

package main

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"
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
	dataCollector := NewHttpDataCollector(forcesEndpoint)
	collector = &ForcesDataProcessor{
		dataCollector: dataCollector,
	}
	return collector
}

func (fdc *ForcesDataProcessor) Process(time time.Time) (forces []Force) {
	forcesData, err := fdc.dataCollector.Collect()
	if err != nil {
		fmt.Printf("Error: %v \n", err)
		return
	}

	forces, err = fdc.Transform(forcesData)
	if err != nil {
		fmt.Printf("Error: %v \n", err)
		return
	}

	wg := sync.WaitGroup{}
	wg.Add(len(forces))

	for force := range forces {
		go func(force int) {
			stopAndSearchProcessor := NewStopAndSearchProcessorFromTimeAndForceId(time, forces[force].Id)
			forces[force].StopAndSearch = stopAndSearchProcessor.GetStopAndSearches()
			wg.Done()
		}(force)
	}

	wg.Wait()
	return
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
