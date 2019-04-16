package cli

import (
	"fmt"

	s3 "github.com/gustavobelfort/42-moulinhate/src/aws-s3"
)

var Status ExamStatus
var fileUrl = "https://s3.amazonaws.com/42-piscine-c/"
var exerciseNames = []string{
	"aff_a",
	"aff_first_param",
	"aff_last_param",
	"aff_z",
	"maff_alpha",
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
		Prefix.LivePrefix = "EXAM > "
		Prefix.IsEnable = true
		GetRandomExercise(getRandomNumber())
	} else {
		fmt.Printf("Error: You have alredy started your exam, finish the current one before running this command.\n")
	}
}

func getRandomNumber() int {
	switch Status.CurrentLevel {
	case 0:
		return (random(0, 7))
	case 1:
		return (random(8, 21))
	case 2:
		return (random(22, 36))
	case 3:
		return (random(37, 51))
	case 4:
		return (random(52, 60))
	case 5:
		return (random(61, 67))
	}
	return 0
}

func GetRandomExercise(number int) {
	s3Folder := fmt.Sprintf("%s%d%s", "e", number, "/subject.en.txt")
	Status.CurrentExercise = s3Folder
	CreateDirectory("./exam/rendu/" + exerciseNames[number])
	CreateDirectory("./exam/subjects/" + exerciseNames[number])
	s3.DownloadFromS3Bucket("./exam/subjects/"+exerciseNames[number]+"/subject.en.txt", fileUrl+s3Folder)
}
