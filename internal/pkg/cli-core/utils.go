package cli

import (
	"fmt"
	"math/rand"
	"os"
)

// CreateDirectory creates the direstories passed to it as parameter
func CreateDirectory(directoryPaths []string) {

	for _, directory := range directoryPaths {
		//choose your permissions well
		pathErr := os.MkdirAll(directoryPath, 0777)

		//check if you need to panic, fallback or report
		if pathErr != nil {
			fmt.Println(pathErr)
		}
	}

}

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}

func getRandomNumber() int {
	switch status.CurrentLevel {
	case 0:
		return (random(0, 11))
	}
	return 0
}
