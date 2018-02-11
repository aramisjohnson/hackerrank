package main
import "fmt"
import "os"
import "strings"
import "strconv"
import (
	"bufio"
	"math"
)

func main() {
	primaryDiagonalTotal := 0
	secondaryDiagonalTotal := 0

	buf := bufio.NewReader(os.Stdin)

	line1, _ := buf.ReadString('\n')
	rowCount, _ := strconv.Atoi(strings.TrimSpace(line1))
	columnIndex := rowCount -1

	for row := 0; row < rowCount; row++ {
		line, _ := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		columns := strings.Split(line, " ")

		priDiagnolValue, _ := strconv.Atoi(columns[row])
		secDiagnolValue, _ := strconv.Atoi(columns[columnIndex - row])

		primaryDiagonalTotal += priDiagnolValue
		secondaryDiagonalTotal += secDiagnolValue
	}

	fmt.Printf("%v", math.Abs(float64(primaryDiagonalTotal) - float64(secondaryDiagonalTotal)))
}
