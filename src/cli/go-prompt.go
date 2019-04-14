package cli

import (
	"strings"

	"github.com/c-bata/go-prompt"
)

var Prefix LivePrefixState

func ChangeLivePrefix() (string, bool) {
	return Prefix.LivePrefix, Prefix.IsEnable
}

func Executor(in string) {
	in = strings.TrimSpace(in)

	blocks := strings.Split(in, " ")
	switch blocks[0] {
	case "start":
		StartExam()
	}
}

func Completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "start", Description: "Start the exam test"},
		{Text: "grademe", Description: "Send your exercise to be graded"},
		{Text: "finish", Description: "Finish your exam and receive your grade"},
		{Text: "help", Description: "See the possiblem commands"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}
