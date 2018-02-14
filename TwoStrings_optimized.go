package main


import "fmt"
import "os"
import (
	"bufio"
	"strings"
	"strconv"
)

func main() {

	buf := bufio.NewReader(os.Stdin)

	line, _ := buf.ReadString('\n')
	line = strings.TrimSpace(line)
	testTotal, _ := strconv.Atoi(line)

	for testCount := 0; testCount < testTotal; testCount++ {
		searchedLetters1 := make(map[string]int)
		searchedLetters2 := make(map[string]int)		
		found := false
		string1, _ := buf.ReadString('\n')
		string1 = strings.TrimSpace(string1)
		stringArray1 := []rune(string1)

		string2, _ := buf.ReadString('\n')
		string2 = strings.TrimSpace(string2)
		stringArray2 := []rune(string2)


		for index1 := 0; index1 < len(stringArray1); index1++ {
			ch1 := string(stringArray1[index1])

			if len(searchedLetters1) == 26 {
				found = false
				break
			}

			//dont search 2nd string if the letter has already been searched once before
			if _, ok := searchedLetters1[ch1]; ok {
				continue
			} else {
				//add to 1st string map if it has not been added already
				searchedLetters1[ch1] = 1
			}

			//search 2nd string map for previous search, if found we are done
			if _, ok := searchedLetters2[ch1]; ok {
				found = true
				break
			}

			for index2 := 0; index2 < len(stringArray2); index2++ {
				ch2 := string(stringArray2[index2])

				//add to 2nd string map if it is not already added
				if _, ok := searchedLetters2[ch2]; !ok {
					searchedLetters2[ch2] = 1
				}

				if strings.EqualFold(ch2, ch1) {
					found = true
					break
				}
			}

			if found {
				break
			}
		}

		if found {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}


	}

}

