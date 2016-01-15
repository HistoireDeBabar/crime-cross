package main

import (
	"fmt"
	"sync"
	"time"
)

type PoliceDataProcessor struct {
	updatedAt                time.Time
	policeForceDataCollector DataCollector
	policeForceTransformer   ForceTransformer
}

func NewDefaultPoliceDataProcessor(time time.Time) DataProcessor {
	policeForceDataCollector := NewHttpDataCollector(forcesEndpoint)
	return &PoliceDataProcessor{
		updatedAt:                time,
		policeForceTransformer:   &ForceDataTransformer{},
		policeForceDataCollector: policeForceDataCollector,
	}
}

func (pdp *PoliceDataProcessor) Process() {
	forcesData, err := pdp.policeForceDataCollector.Collect()
	if err != nil {
		fmt.Printf("Error: %v \n", err)
		return
	}

	forces, err := pdp.policeForceTransformer.Transform(forcesData)
	if err != nil {
		fmt.Printf("Error: %v \n", err)
		return
	}

	wg := sync.WaitGroup{}
	wg.Add(len(forces))

	for force := range forces {
		go func(force int) {
			stopAndSearchProcessor := NewStopAndSearchProcessorFromTimeAndForceId(pdp.updatedAt, forces[force].Id)
			forces[force].StopAndSearch = stopAndSearchProcessor.GetStopAndSearches()
			wg.Done()
		}(force)
	}

	wg.Wait()
	return
}
