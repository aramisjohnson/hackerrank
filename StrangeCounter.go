package main


import "fmt"
import "strings"
import "os"
import (
	"bufio"
	"strconv"
)

func main() {
	buf := bufio.NewReader(os.Stdin)

	line1, _ := buf.ReadString('\n')
	timer, _ := strconv.Atoi(strings.TrimSpace(line1))
	lowerRange := 1

	for {
		value := lowerRange + 2 
		upperRange := lowerRange + value

		if timer >= lowerRange && timer < upperRange {
			fmt.Printf("%v\n", lowerRange - timer + value)
			break
		}

		lowerRange = upperRange
	}
}
