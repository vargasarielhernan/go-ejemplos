package main

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

// create a new scheduler
func main() {
	loc, err := time.LoadLocation("America/Argentina/Buenos_Aires")
	if err != nil {
		panic(err)
	}
	s := gocron.NewScheduler(loc)
	s.Every(1).Second().Do(task)
	s.StartBlocking()
}
func task() {
	fmt.Println("hello")
}
