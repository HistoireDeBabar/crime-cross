package main

type DataCollector interface {
	Collect() ([]byte, error)
}
