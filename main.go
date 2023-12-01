package main

import (
	"go-calculator/application"
	"sync"
)

func main() {
	// Create a new instance of the application.
	app := application.New()

	// Create a wait group to manage the goroutine
	// This is not necessary for this, just for practice
	wg := sync.WaitGroup{}

	// Run the application.
	wg.Add(1)
	go app.Run(&wg)

	// Wait for the goroutine to finish
	wg.Wait()
}