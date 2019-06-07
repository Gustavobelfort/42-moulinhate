package download

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/Gustavobelfort/42-moulinhate/internal/pkg/status"
)

var fileURL = "https://gustavobelfort.dev/42-moulinhate/exercises/"

// Exercise fetches the assigned exercise number from the server
func Exercise(exerciseNumber, exerciseLevel int) {
	metadata := metadataFromServer(exerciseNumber, exerciseLevel)
	subjectFromServer(metadata)
}

func metadataFromServer(exerciseNumber, exerciseLevel int) status.Exercise {

	var info status.Exercise

	metadataPath := fmt.Sprintf("%s/%d/%d/metadata.json", fileURL, exerciseLevel, exerciseNumber)

	resp, err := http.Get(metadataPath)
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}

	output := json.Unmarshal([]byte(body), &info)
	if output != nil {
		panic(output.Error())
	}
	return info
}

func subjectFromServer(metadata status.Exercise) {

	subjectPath := fmt.Sprintf("%s/%d/%d/subject.en.txt", fileURL, metadata.ExerciseLevel, metadata.ExerciseNumber)

	resp, err := http.Get(subjectPath)
	if err != nil {
		panic(err.Error())
	}

	out, err := os.Create("./exam/subjects/" + metadata.ExerciseName + "/subject.en.txt")
	if err != nil {
		panic(err.Error())
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		panic(err.Error())
	}

	pathErr := os.MkdirAll("./exam/rendu/"+metadata.ExerciseName, 0777)
	if pathErr != nil {
		fmt.Println(pathErr)
	}
}

func testcasesFromServer(metadata status.Exercise) ([][]string, error) {

	testcasesPath := fmt.Sprintf("%s/%d/%d/testcases.csv", fileURL, metadata.ExerciseLevel, metadata.ExerciseNumber)

	resp, err := http.Get(testcasesPath)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	reader := csv.NewReader(resp.Body)
	reader.Comma = ';'

	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}
