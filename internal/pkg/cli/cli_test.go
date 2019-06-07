package cli_test

import (
	"os"
	"testing"

	"github.com/Gustavobelfort/42-moulinhate/internal/pkg/cli"
	"github.com/Gustavobelfort/42-moulinhate/internal/pkg/status"
)

func TestStartExam(t *testing.T) {

	testExamStatus := status.Exam{
		CurrentLevel:    0,
		RemainingLives:  3,
		Started:         false,
		CurrentExercise: "",
		Retries:         2,
	}

	folders := []string{
		"./exam/rendu",
		"./exam/subjects",
		"./exam/traces",
	}

	t.Run("Running when the test has not started", func(t *testing.T) {
		cli.StartExam(testExamStatus)

		for _, directory := range folders {

			if _, err := os.Stat(directory); os.IsNotExist(err) {
				// Directory does *not* exist
				t.Errorf("%s wasn't created properly", directory)
			}
		}

		os.RemoveAll("./exam")

	})

	t.Run("Running when the test has already started", func(t *testing.T) {

		testExamStatus.Started = true
		if err := cli.StartExam(testExamStatus); err == nil {
			t.Errorf("The function did not returned error if the test had already started")
		}

		os.RemoveAll("./exam")
	})

}
