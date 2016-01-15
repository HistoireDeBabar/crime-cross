package main

import (
	"fmt"
	"sync"
)

func main() {
	dateChecker := NewDefaultUpdateChecker()
	result, update := dateChecker.Check()
	fmt.Printf("Can Update: %v \n", result)

	if !result {
		DestoryStack()
		return
	}
	dataProcessors := []DataProcessor{
		NewDefaultPoliceDataProcessor(update),
	}
	Process(dataProcessors)
	fmt.Println("Data Shipped")
}

func DestoryStack() {
	// Send an event to lambda to destroy the stack.
}

// Process a list of DataProcessors.
func Process(dataProcessors []DataProcessor) {
	wg := sync.WaitGroup{}
	wg.Add(len(dataProcessors))
	for process := range dataProcessors {
		go func(process int) {
			dataProcessors[process].Process()
			wg.Done()
		}(process)
	}

	wg.Wait()
	return
}
