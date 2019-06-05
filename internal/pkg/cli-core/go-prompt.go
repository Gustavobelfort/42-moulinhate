package cli

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/Gustavobelfort/42-moulinhate/pkg/correct-exercise"
	"github.com/c-bata/go-prompt"
	"github.com/common-nighthawk/go-figure"
)

var prefix LivePrefixState

func changeLivePrefix() (string, bool) {
	return prefix.LivePrefix, Prefix.IsEnable
}

func executor(in string) {
	in = strings.TrimSpace(in)

	blocks := strings.Split(in, " ")
	switch blocks[0] {
	case "start":
		StartExam()
	case "grademe":
		correct.CorrectExercise()
	}
}

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "start", Description: "Start the exam test"},
		{Text: "grademe", Description: "Send your exercise to be graded"},
		{Text: "finish", Description: "Finish your exam and receive your grade"},
		{Text: "help", Description: "See the possiblem commands"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

// Start runs the go-prompt cli
func Start() {
	rand.Seed(time.Now().UnixNano())
	myFigure := figure.NewFigure("42 - MoulinHate", "", true)
	myFigure.Print()
	p := prompt.New(
		cli.Executor,
		cli.Completer,
		prompt.OptionLivePrefix(cli.ChangeLivePrefix),
		prompt.OptionPrefix("piscine-examshell >"),
		prompt.OptionTitle("42 - MoulinHate"),
	)
	fmt.Printf(`-----------------------------------------------------------------------------------
Welcome to the 42-MoulinHate exam simulator for the piscine exams!
type help to see the available commands or type start to start a exam simulation
Good Luck ;)
`)
	p.Run()
}
