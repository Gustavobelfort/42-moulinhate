package main

import (
	"fmt"
	"math/rand"
	"time"

	"42-moulinhate/internal/pkg/cli"

	"github.com/c-bata/go-prompt"
	"github.com/common-nighthawk/go-figure"
)

func main() {
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