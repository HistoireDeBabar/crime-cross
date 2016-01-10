package main

import (
	"fmt"
	"testing"
)

func TestS3Connection(t *testing.T) {
	s3Data := NewDefaultS3DataCollector()
	data, err := s3Data.Collect()
	fmt.Println(data)
	fmt.Println(err)
}
