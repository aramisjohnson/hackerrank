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
	minValue := 0
	maxValue := 0

	buf := bufio.NewReader(os.Stdin)

	line1, _ := buf.ReadString('\n')
	inputs := strings.Split(strings.TrimSpace(line1), " ")
	length := len(inputs)

	//sort inputs
	for swapped {
		swapped = false

		for index := 0; index < length-1; index++ {
			val1, _ := strconv.Atoi((inputs[index]))
			val2, _ := strconv.Atoi((inputs[index+1]))

			if  val1 > val2 {
				inputs[index] = strconv.Itoa(val2)
				inputs[index+1] = strconv.Itoa(val1)
				swapped = true
			}
		}
	}


	for index := 0; index < length-1; index++ {
		val, _ := strconv.Atoi(inputs[index])
		minValue += val

		val, _ = strconv.Atoi(inputs[length - index -1])
		maxValue += val
	}

	fmt.Printf("%v %v", minValue, maxValue)
}
