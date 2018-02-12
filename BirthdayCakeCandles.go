package main


import "fmt"
import "strings"
import "strconv"
import "os"
import (
	"bufio"
)

func main() {
	candleMap := make(map[int]int)
	candleKey := 0

	buf := bufio.NewReader(os.Stdin)

	line1, _ := buf.ReadString('\n')
	candleCount, _ := strconv.Atoi(strings.TrimSpace(line1))
	line2, _ := buf.ReadString('\n')
	inputs := strings.Split(strings.TrimSpace(line2), " ")


	for index := 0; index < candleCount; index++ {
		candleHeight, _ := strconv.Atoi((inputs[index]))

		if val, ok := candleMap[candleHeight]; ok {
			candleMap[candleHeight] = val + 1
		} else {
			candleMap[candleHeight] = 1
		}

		if candleKey < candleHeight {
			candleKey = candleHeight
		}
	}

	fmt.Printf("%v",candleMap[candleKey])
}
