package cli

import (
	"fmt"

	download "github.com/Gustavobelfort/42-moulinhate/pkg/download-exercise"
)

var folders = []string{
	"./exam/rendu",
	"./exam/subjects",
	"./exam/traces",
}

func StartExam() {
	if !status.Started {
		status.Started = true
		CreateDirectory(folders)
		Prefix.LivePrefix = "EXAM > "
		Prefix.IsEnable = true
		GetRandomExercise(GetRandomExerciseNumber())
	} else {
		fmt.Printf("Error: You have alredy started your exam, finish the current one before running this command.\n")
	}
}

func startRandomExercise(number int) {
	CreateDirectory("./exam/rendu/" + ExerciseNames[number])
	CreateDirectory("./exam/subjects/" + ExerciseNames[number])
	download.FromServer(number)
}
