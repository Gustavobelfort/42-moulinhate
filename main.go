package main

import (
	"fmt"

	"github.com/gustavobelfort/42-moulinhate/src/cli"

	"github.com/c-bata/go-prompt"
	"github.com/common-nighthawk/go-figure"
)

// func getRandomExercise(level, retires int) {
// 	Dir()
// }

func main() {
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
