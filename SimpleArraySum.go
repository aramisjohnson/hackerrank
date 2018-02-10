package main
import "fmt"
import "os"
import "strings"
import "strconv"
import "bufio"

func main() {
    buf := bufio.NewReader(os.Stdin)

    line1, _ := buf.ReadString('\n')
    line1 = strings.TrimSpace(line1)
	  line2, _ := buf.ReadString('\n')
    line2 = strings.TrimSpace(line2)
    
    length, _ := strconv.Atoi(line1)

    inputs := strings.Split(line2, " ")
    
    total := 0
    
    for x := 0; x < length; x++ {
        input, _ := strconv.Atoi(inputs[x])
        
        total += input
    }
    
    fmt.Printf("%v", total)
}
