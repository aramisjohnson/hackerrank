package main
import "fmt"
import "strings"
import "strconv"
import "os"
import "bufio"

func main() {
	buf := bufio.NewReader(os.Stdin)

	line1, _ := buf.ReadString('\n')
	alice := strings.Split(strings.TrimSpace(line1), " ")
	line2, _ := buf.ReadString('\n')
	bob := strings.Split(strings.TrimSpace(line2), " ")

	alicePoints := 0
	bobPoints := 0

	for i := 0; i < len(alice); i++ {
		aliceScore, _ := strconv.Atoi(alice[i])
        bobScore, _ := strconv.Atoi(bob[i])

        if aliceScore == bobScore {
            continue
        } else if aliceScore > bobScore {
			alicePoints += 1
		} else {
			bobPoints += 1
		}
	}

	fmt.Printf("%v %v", alicePoints, bobPoints)
}
