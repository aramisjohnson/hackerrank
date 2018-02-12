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

	startIndex := 3
	internalIndex := startIndex

	for externalTimer := 1; ; externalTimer++ {

		if externalTimer == timer {
			break
		} else if internalIndex == 1 {
			startIndex += startIndex
			internalIndex = startIndex
		} else {
			internalIndex--
		}
	}

	fmt.Printf("%v\n", internalIndex)
}
