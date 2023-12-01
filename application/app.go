package application

import "fmt"

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
func (app *App) Run() {
	
	// welcome user
	app.Welcome()

}

// welcome function
func (app *App) Welcome() {
	
	fmt.Println("\nWelcome to " + app.name)
	fmt.Println("Solve all your maths problems here!")
	fmt.Println("\n=================================")
	fmt.Println()

}