//Package check will handle determining if the user won or not.
package check

import "fmt"

//GotWrong is the amount the user guessed incorrectly.
var GotWrong int8 = 0

//HowMany checks the amount you got wrong.
func HowMany() {

	//The user lost if GotWrong > 3. If GotWrong <= 3, the user won.
	switch {

	case GotWrong > 3:
		fmt.Println("You got", GotWrong, "wrong. You lost.")

	default:
		fmt.Println("You got", GotWrong, "wrong. You won!")

	}

}
