package main


import "fmt"
import "os"
import (
	"bufio"
	"strings"
	"unicode"
	"math"
)

func main() {
	neededDeletionCount := int64(0)


	buf := bufio.NewReader(os.Stdin)

	line1, _ := buf.ReadString('\n')
	line1 = strings.TrimSpace(line1)
	line2, _ := buf.ReadString('\n')
	line2 = strings.TrimSpace(line2)

	wordMap1 := fillMapWithDefaults()
	wordMap2 := fillMapWithDefaults()

	wordMap1 = addToMap(wordMap1, line1)
	wordMap2 = addToMap(wordMap2, line2)

	for index := 1; index <= 26; index++ {
		letter := string( unicode.ToLower(rune('A' -1 + index)))

		neededDeletionCount += int64(math.Abs (float64(wordMap1[letter] - wordMap2[letter])))
	}


	fmt.Printf("%v\n", neededDeletionCount)

}

func fillMapWithDefaults() (letterMap map[string]int) {
	letterMap = make(map[string]int)
	for index := 1; index <=26; index++ {
		letterMap[string(unicode.ToLower(rune('A' -1 + index)))] = 0
	}

	return
}

func addToMap(letterMap map[string]int, word string) (map[string]int) {
	runes := []rune(word)

	for _, c := range runes {
		if val, ok := letterMap[string(c)]; ok {
			letterMap[string(c)] = val + 1
		} else {
			letterMap[string(c)] = 1
		}
	}

	return letterMap
}
