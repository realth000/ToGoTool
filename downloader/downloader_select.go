package downloader

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// GetSelectFromUser shows all choices in chooseMap, let user selects one and returns the select value.
func GetSelectFromUser(chooseMap map[string]interface{}) interface{} {
	// Get print format
	maxLength := 0
	var choices []string // Used to access chooseMap map in decreasing order.
	for choice := range chooseMap {
		if maxLength < len(choice) {
			maxLength = len(choice)
		}
		choices = append(choices, choice)
	}

	// Sort choices in decreasing order.
	sort.Sort(sort.Reverse(sort.StringSlice(choices)))

	// Print
	index := 0
	fmt.Printf("Index  \n")
	for _, name := range choices {
		//fmt.Printf(" %2d    %-"+strconv.Itoa(maxLength)+"s  %s\n", index, name, chooseMap[name])
		fmt.Printf(" %2d    %-"+strconv.Itoa(maxLength)+"s\n", index, name)
		index++
	}

	// Receive index to download from standard input.
	chooseFunc := func() interface{} {
		for {
			reader := bufio.NewReader(os.Stdin)
			fmt.Printf("Choose one index to download[0, %d]: ", index-1)
			i, _ := reader.ReadString('\n')
			j, err := strconv.Atoi(i[:len(i)-1])
			if err != nil {
				fmt.Println("Invalid index")
				continue
			}
			if j < 0 || j >= len(chooseMap) {
				fmt.Println("Index out of range")
				continue
			}
			fmt.Printf("Download index=%d\n", j)
			return chooseMap[choices[j]]
		}
	}()
	return chooseFunc
}
