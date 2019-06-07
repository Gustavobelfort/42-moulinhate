package cli

import (
	"fmt"

	// download "github.com/Gustavobelfort/42-moulinhate/internal/pkg/download"
	status "github.com/Gustavobelfort/42-moulinhate/internal/pkg/status"
)

var folders = []string{
	"./exam/rendu",
	"./exam/subjects",
	"./exam/traces",
}

func StartExam(exam status.Exam) error {
	
	if !exam.Started {
		exam.Started = true
		CreateDirectory(folders)
		// download.Exercise(getRandomExerciseNumber(exam), exam.CurrentLevel)
	} else {
		return fmt.Errorf("error: You have alredy started your exam, finish the current one before running this command")
	}

	return nil
}
