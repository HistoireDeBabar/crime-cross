package main

import (
	"time"
)

type Checker interface {
	Check() (bool, time.Time)
}
