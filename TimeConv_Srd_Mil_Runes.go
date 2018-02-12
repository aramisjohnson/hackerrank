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
	length := len(line1)
	stdTimeString := strings.TrimSpace(line1)

	runes := []rune(stdTimeString)

	timeSubstring := string(runes[0:length-2])

	timeArray := strings.Split(timeSubstring, ":")

	hr, _ := strconv.Atoi(timeArray[0])

    periodString := string(runes[len(runes)-2:])

	if strings.EqualFold(periodString, "PM") && hr != 12 {
		hr += 12
	}

	if strings.EqualFold(periodString, "AM") && hr == 12 {
		hr = 0
	}

	fmt.Printf("%02d:%v:%v\n", hr ,timeArray[1], timeArray[2])
}
