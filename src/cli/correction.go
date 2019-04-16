package cli

import (
	"bytes"
	"fmt"
	"os/exec"
)

func Correct() {

	command := exec.Command("bash", "run "+Status.CurrentExercise)
	var out bytes.Buffer

	command.Stdout = &out
	err := command.Run()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(out.String())
}
