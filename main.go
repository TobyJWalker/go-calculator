package main

import (
	"go-calculator/application"
)

func main() {
	// 1. Create a new instance of the application.
	app := application.New()

	// 2. Run the application.
	app.Run()
}