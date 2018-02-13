package main


import "fmt"
import "os"
import (
	"bufio"
	"strings"
)

func main() {
	buf := bufio.NewReader(os.Stdin)
	numberOfScragullers := 0
	line, _ := buf.ReadString('\n')
	line = strings.TrimSpace(line)


	//store line in an array
	arrOfString := []rune(line)
	length := len(arrOfString)

	characterMap := make(map[string]int)

	//add characters to a map
	for index := 0; index < length; index++ {
		ch := string(arrOfString[index])
		if val, ok := characterMap[ch]; ok && val != 2{
			characterMap[ch] = val + 1
		} else {
			characterMap[ch] = 1
		}
	}

	for _, v := range characterMap {
		if length % 2 > 0 && v == 1 {
			numberOfScragullers++
		} else if v % 2 != 0 {
			numberOfScragullers = 100
			break
		}
	}

	if numberOfScragullers > 1 {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}

}

