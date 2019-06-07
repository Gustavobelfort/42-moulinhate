package cli

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"os"

	status "github.com/Gustavobelfort/42-moulinhate/internal/pkg/status"
)

// CreateDirectory creates the direstories passed to it as parameter
func CreateDirectory(directoryPaths []string) {

	for _, directory := range directoryPaths {
		//choose your permissions well
		pathErr := os.MkdirAll(directory, 0777)
		//check if you need to panic, fallback or report
		if pathErr != nil {
			fmt.Println(pathErr)
		}
	}

}

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}

func getRandomExerciseNumber(exam status.Exam) int {
	switch exam.CurrentLevel {
	case 0:
		return (random(0, 11))
	}
	return 0
}

func CaptureOutput(f func()) string {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	f()
	log.SetOutput(os.Stderr)
	return buf.String()
}
