package cli

import (
	"fmt"
	"math/rand"
	"os"
)

func CreateDirectory(directoryPath string) {
	//choose your permissions well
	pathErr := os.MkdirAll(directoryPath, 0777)

	//check if you need to panic, fallback or report
	if pathErr != nil {
		fmt.Println(pathErr)
	}
}

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}
