package main


import "fmt"
import "strings"
import "os"
import (
	"bufio"
	"strconv"
	"math"
)

func main() {
	buf := bufio.NewReader(os.Stdin)

	line1, _ := buf.ReadString('\n')
	testCount, _ := strconv.Atoi(strings.TrimSpace(line1))


	for testIndex := 0; testIndex < testCount; testIndex++ {
		line, _ := buf.ReadString('\n')
		lineArray := strings.Split(strings.TrimSpace(line), " ")
		N, _ := strconv.ParseInt(lineArray[0], 10, 64)
		K, _ := strconv.ParseInt(lineArray[1], 10, 64)

		numArray := make([]int64, N)

		//fill array
		for count := int64(1); count <= N; count++ {
			numArray[count-1] = count
		}

		err := doPermute(numArray, 0, K, N-1)

		if err != nil {
			fmt.Printf("%v\n", err.Error())
		}
	}
}

func doPermute(inArray []int64, level int64, K int64, N int64) (err error) {
	err = TestForK(inArray, K)
	if err == nil {
		return err
	}

	for index := level; index <= N; index++ {
		inArray = swap(inArray, level, index)

		err = doPermute(inArray, level+1, K, N)

		if err == nil {
			return
		}

		inArray = swap(inArray, level, index)
	}

	return fmt.Errorf("-1")
}

func swap(inArray []int64, index1 int64, index2 int64) (outArray []int64){
	if index1 == index2 {return inArray}
	(inArray)[index1], (inArray)[index2] = (inArray)[index2], (inArray)[index1]

	return inArray
}

func TestForK(inArray []int64, K int64) (err error) {
	//Test for K
	for index := int64(1); index <= int64(len(inArray)); index++ {
		absDiff := int64(math.Abs(float64(index - inArray[index-1])))
		if absDiff != K {
			return fmt.Errorf("-1")
		}
	}

	for index := int64(0); index <= int64(len(inArray))-1; index++ {
		fmt.Printf("%v ", inArray[index])
	}

	fmt.Printf("\n")

	return nil
}
