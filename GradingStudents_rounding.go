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
	line1 = strings.TrimSpace(line1)

	numOfGrades, _ := strconv.Atoi(line1)


	for count := 0; count < numOfGrades; count++ {
		line, _ := buf.ReadString('\n')
		line = strings.TrimSpace(line)

		grade, _ := strconv.Atoi(line)

		if grade < 38 {
			fmt.Printf("%v\n", grade)
			continue
		}

		floatGrade := float64(grade)

		roundedGrade := int(float64(floatGrade / 5.0) + .5) * 5
		
		if grade > roundedGrade {
			roundedGrade = grade
		}
		
		fmt.Printf("%v\n", int(roundedGrade))
	}
}
