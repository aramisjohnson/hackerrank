package main


import "fmt"
import "strings"
import "strconv"
import "os"
import (
	"bufio"
)

func main() {
	buf := bufio.NewReader(os.Stdin)

	line1, _ := buf.ReadString('\n')
	floors, _ := strconv.Atoi(strings.TrimSpace(line1)) 

	for index := 1; index <= floors; index++ {
		fmt.Printf("%v%v\n", strings.Repeat(" ", floors - index), strings.Repeat("#", index))
	}
}
