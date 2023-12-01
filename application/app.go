package application

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

// create app struct
type App struct {
	name string
}

// construct app
func New() *App {
	return &App{
		name: "Go Calculator",
	}
}

// run function
func (app *App) Run(wg *sync.WaitGroup) {
	
	// welcome user
	app.Welcome()

	// loop indefinitely
	for {

		// get user input
		input := app.GetUserInput()
		fmt.Println(input)

	}

	// mark done
	wg.Done()
}

// welcome function
func (app *App) Welcome() {
	
	fmt.Println("\nWelcome to " + app.name)
	fmt.Println("Solve all your maths problems here!")
	fmt.Println("\n=================================")
	fmt.Println()

}

// get user input
func (app *App) GetUserInput() string {

	// initialise bufio scanner
	scanner := bufio.NewScanner(os.Stdin)

	// prompt user
	fmt.Print(">> ")
	scanner.Scan()
	input := scanner.Text()

	return input

}

