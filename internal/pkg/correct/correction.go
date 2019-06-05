package correct

import (
	"fmt"
	"log"
	"os/exec"
	
	status "github.com/Gustavobelfort/42-moulinhate/internal/pkg/status"
)

// Exercise looks for the .c submmited as answer in the path of the exercise and call the assigned check functions
func Exercise(status status.Exam) {

	out, err := exec.Command("gcc -Wall -Werror -Wextra -o " + status.CurrentExercise + "").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The date is %s\n", out)
}
