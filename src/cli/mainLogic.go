package cli

import (
	"fmt"
	"math/rand"

	s3 "github.com/gustavobelfort/42-moulinhate/src/aws-s3"
)

var Status ExamStatus
var fileUrl = "https://s3.amazonaws.com/42-piscine-c/"
var exerciseNames = []string{
	"aff_a",
	"aff_first_param",
	"aff_last_param",
	"aff_z", "maff_alpha",
	"maff_revalpha",
	"only_z",
	"strlen",
}

func StartExam() {

	if !Status.Started {
		Status.Started = true
		CreateDirectory("./exam/rendu")
		CreateDirectory("./exam/subjects")
		CreateDirectory("./exam/traces")
		Prefix.LivePrefix = "EXAM " + " > "
		Prefix.IsEnable = true
		GetRandomExercise(getRandomNumber())
	} else {
		fmt.Printf("Error: You have alredy started your exam, finish the current one before running this command.\n")
	}
}

func getRandomNumber() int {
	switch Status.CurrentLevel {
	case 0:
		min := 0
		max := 7
		return (rand.Intn(max-min) + min)
	case 1:
		min := 8
		max := 21
		return (rand.Intn(max-min) + min)
	case 2:
		min := 22
		max := 36
		return (rand.Intn(max-min) + min)
	case 3:
		min := 37
		max := 51
		return (rand.Intn(max-min) + min)
	case 4:
		min := 52
		max := 60
		return (rand.Intn(max-min) + min)
	case 5:
		min := 61
		max := 67
		return (rand.Intn(max-min) + min)
	}
	return 0
}

func GetRandomExercise(number int) {
	s3Folder := fmt.Sprintf("%s%d%s", "e", number, "/subject.en.txt")
	CreateDirectory("./exam/rendu/" + exerciseNames[number])
	CreateDirectory("./exam/subjects/" + exerciseNames[number])
	s3.DownloadFromS3Bucket("./exam/subjects/"+exerciseNames[number]+"/subject.en.txt", fileUrl+s3Folder)
}
