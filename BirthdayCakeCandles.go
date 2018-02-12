package main


import "fmt"
import "strings"
import "strconv"
import "os"
import (
	"bufio"
)

func main() {
	swapped := true
	maxCandleCount := 0
	buf := bufio.NewReader(os.Stdin)

	line1, _ := buf.ReadString('\n')
	candleCount, _ := strconv.Atoi(strings.TrimSpace(line1))
	line2, _ := buf.ReadString('\n')
	inputs := strings.Split(strings.TrimSpace(line2), " ")

	//sort inputs
	for swapped {
		swapped = false

		for index := 0; index < candleCount-1; index++ {
			val1, _ := strconv.Atoi((inputs[index]))
			val2, _ := strconv.Atoi((inputs[index+1]))

			if  val1 > val2 {
				maxCandleCount = 0
				inputs[index] = strconv.Itoa(val2)
				inputs[index+1] = strconv.Itoa(val1)
				swapped = true
			 } else if val1 == val2 {
			 	maxCandleCount++
			}
		}
	}

	fmt.Printf("%v",maxCandleCount)
}
