package main

import (
	//Basic packages
	"fmt"

	//These  packages help me get user input.
	"bufio"
	"os"
	"strconv"
	"strings"

	//The math/rand package allows me to generate random numbers.
	"math/rand"

	//The time package allows me to use time.Now() for a unique rand.Seed()
	"time"

	//My check package from ./check/check.go.
	"github.com/NewbJS/another/check"
)

func main() {

	//Generates a random 64 bit integer from 0 to 10.
	rand.Seed(time.Now().UnixNano())
	ranNum := rand.Int63n(11)
	fmt.Println("Pick a random number, from 0 to 10!")

	//Setting up the scanner
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	guessedCorrectly := false

	//Main game logic. This should continue to loop until the user guessed correctly.
	for !guessedCorrectly {
		//If the input is a number, the type will be int64. If the input is not number, input = 0.
		input, err := strconv.ParseInt(scanner.Text(), 10, 64)

		nan := false

		//Check for an error.
		if err != nil {
			input = 100
			nan = true
			fmt.Println("That is not a number. Try again.")
			check.GotWrong--
		}

		switch {

		//Add one to the check.GotWrong variable if input > ranNum. Scan afterward.
		case input > ranNum:
			if !nan {
				fmt.Println("Guess Lower!")
			}

			check.GotWrong++
			scanner.Scan()

		//Add on to the check.GotWrong variable if input < ranNum. Scan afterward.
		case input < ranNum:
			if !nan {
				fmt.Println("Guess higher!")
			}
			check.GotWrong++
			scanner.Scan()

		//If both previous cases returned false, the user guessed correctly.
		default:

			//End the loop, and tell the user how many guesses they got wrong. check.HowMany() is a function form the check package.

			check.HowMany()
			guessedCorrectly = true
			checkIfPlayAgain()
		}

	}

}

func checkIfPlayAgain() {

	//Scanner configuration
	fmt.Println("Do you want to play again? y/n")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := strings.ToLower(scanner.Text())

	//Restart if the user typed "y". Otherwise, exit.
	if input == "y" {
		check.GotWrong = 0
		main()
	} else {
		fmt.Println("Exited.")

	}
}
