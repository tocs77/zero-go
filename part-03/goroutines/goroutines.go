//--Summary:
//  Create a program to read a list of numbers from multiple files,
//  sum the total of each file, then sum all the totals.
//
//--Requirements:
//* Sum the numbers in each file noted in the main() function
//* Add each sum together to get a grand total for all files
//  - Print the grand total to the terminal
//* Launch a goroutine for each file
//* Report any errors to the terminal
//
//--Notes:
//* This program will need to be ran from the `lectures/exercise/goroutines`
//  directory:
//    cd lectures/exercise/goroutines
//    go run goroutines
//* The grand total for the files is 4103109
//* The data files intentionally contain invalid entries
//* stdlib packages that will come in handy:
//  - strconv: parse the numbers into integers
//  - bufio: read each line in a file
//  - os: open files
//  - io: io.EOF will indicate the end of a file
//  - time: pause the program to wait for the goroutines to finish

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func sumFile(file string) int {
	dat, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	strData := strings.Split(string(dat), "\n")
	sum := 0
	for _, line := range strData {
		num, err := strconv.Atoi(strings.TrimSpace(line))
		if err != nil {
			fmt.Println(err)
			continue
		}
		sum += num
	}
	return sum
}

func main() {
	files := []string{"num1.txt", "num2.txt", "num3.txt", "num4.txt", "num5.txt"}

	var sums []int = make([]int, len(files))

	for i := 0; i < len(files); i++ {
		go func(i int) {
			sums[i] = sumFile(files[i])
		}(i)
	}

	time.Sleep(100 * time.Millisecond)

	sum := 0
	for i := 0; i < len(files); i++ {
		sum += sums[i]
	}
	fmt.Printf("sum = %d\n", sum)
}
