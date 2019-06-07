package correct

import (
	"fmt"
	"log"
	"os/exec"

	status "github.com/Gustavobelfort/42-moulinhate/internal/pkg/status"
)

// Exercise looks for the .c submmited as answer in the path of the exercise and call the assigned check functions
func Exercise(status status.Exam) {

	buildCommand := fmt.Sprintf("%s%s%s%s%s", "gcc -Wall -Werror -Wextra -o ./exam/traces/", status.CurrentExercise, " ./exam/subjects/", status.CurrentExercise, "/*")
	out, err := exec.Command(buildCommand).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The date is %s\n", out)
}
