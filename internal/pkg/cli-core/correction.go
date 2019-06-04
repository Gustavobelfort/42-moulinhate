package cli

import (
	"fmt"
	"log"
	"os/exec"
)

// Correct looks for the .c submmited as answer in the path of the exercise and call the assigned check functions
func Correct() {

	out, err := exec.Command("bash", "run "+status.CurrentExercise).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The date is %s\n", out)
}
