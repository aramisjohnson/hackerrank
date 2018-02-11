package main


import "fmt"
import "strings"
import "strconv"
import "os"
import (
	"bufio"
)

func main() {
	var zeroPercent = 0.000000
	var posPercent = 0.000000
	var negPercent = 0.000000

	buf := bufio.NewReader(os.Stdin)

	line1, _ := buf.ReadString('\n')
	total, _ := strconv.Atoi(strings.TrimSpace(line1))
	line2, _ := buf.ReadString('\n')
	input := strings.Split(strings.TrimSpace(line2), " ")

	positiveCount := 0
	negativeCount := 0
	zeroCount := 0

	for i := 0; i < len(input); i++ {
		val, _ := strconv.Atoi(input[i])
		if val < 0 {
			negativeCount++
		} else if val > 0 {
			positiveCount++
		} else {
			zeroCount++
		}
	}

	if zeroCount != 0 {
		zeroPercent = float64(zeroCount) / float64(total)
	}

	if positiveCount != 0 {
		posPercent = float64(positiveCount)/ float64(total)
	}

	if negativeCount != 0 {
		negPercent = float64(negativeCount) / float64(total)
	}

	fmt.Printf("%0.6f\n", posPercent)
	fmt.Printf("%0.6f\n", negPercent)
	fmt.Printf("%0.6f\n",zeroPercent)
}
