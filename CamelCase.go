package main


import "fmt"
import "os"
import (
	"bufio"
	"unicode"
	"strings"
)

func main() {
	wordCount := 0
	buf := bufio.NewReader(os.Stdin)

	line1, _ := buf.ReadString('\n')
	line1 = strings.TrimSpace(line1)

	if !strings.EqualFold(line1, "") {
		wordCount++
	}

	stringSequence := []rune(line1)

	for _,c := range stringSequence{
		if unicode.IsUpper(c) {
			wordCount++
		}
	}

	fmt.Printf("%v\n", wordCount)

}

