package main

import (
	"fmt"
	"time"
)

func main() {
	subject := NewDefaultUpdateChecker()
	result, update := subject.Check()
	fmt.Printf("Can Update: %v \n", result)

	if !result {
		DestoryStack()
		return
	}
	ProcessPoliceData(update)
}

func DestoryStack() {

}

func ProcessPoliceData(time time.Time) {
	forcesProcessor := NewDefaultForcesDataProcessor()
	forces := forcesProcessor.Process(time)
	fmt.Println(forces)
}
