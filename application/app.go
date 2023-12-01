package application

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

// create app struct
type App struct {
	name string
}

// equation structure
type Equation struct {
	num1 float64
	num2 float64
	op string
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
		input, exit_sig := app.GetUserInput()

		// check for exit signal
		if exit_sig {
			break
		}
		
		// parse user input and handle errors
		eq, err := app.ParseInput(input)
		if err {
			continue
		}

		//display equation
		fmt.Println("Equation: ", eq.num1, eq.op, eq.num2)

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
func (app *App) GetUserInput() (string, bool) {

	// initialise bufio scanner
	scanner := bufio.NewScanner(os.Stdin)

	// prompt user
	fmt.Print(">> ")
	scanner.Scan()
	input := scanner.Text()

	// check for exit request
	if strings.Contains(input, "exit") {
		return input, true
	} else {
		return input, false
	}

}

// parse user input
func (app *App) ParseInput(input string) (*Equation, bool) {

	// split input string
	input_arr := strings.Split(input, " ")

	// check for valid input
	if len(input_arr) != 3 {
		fmt.Println("Invalid input")
		return nil, true
	}

	// extract values
	num1, err_1 := strconv.ParseFloat(input_arr[0], 64) // float64
	op := input_arr[1]
	num2, err_2 := strconv.ParseFloat(input_arr[2], 64) // float64

	// check for valid operator
	if op != "+" && op != "-" && op != "*" && op != "/" {
		fmt.Println("Invalid operator")
		return nil, true
	}

	// check for valid numbers
	if err_1 != nil || err_2 != nil {
		fmt.Println("Invalid numbers")
		return nil, true
	}

	// create equation
	eq := &Equation{
		num1: num1,
		num2: num2,
		op: op,
	}

	// return equation, no errors
	return eq, false

}