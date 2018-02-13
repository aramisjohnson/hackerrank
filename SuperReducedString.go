package main


import "fmt"
import "os"
import (
	"bufio"
)

func main() {
	changes := true
	buf := bufio.NewReader(os.Stdin)

	line1, _ := buf.ReadString('\n')

	stringSequence := []rune(line1)

	for changes {
		changes = false

		for index := 0; index < len(stringSequence)-1; index++ {
			if stringSequence[index] == stringSequence[index+1] {
				//remove it
				stringSequence = append(stringSequence[:index], stringSequence[index+2:]...)
				changes = true
				break
			}
		}
	}

	if len(stringSequence) == 0 {
		fmt.Println("Empty String")
	} else {
		fmt.Printf("%v\n", string(stringSequence))
	}

}

