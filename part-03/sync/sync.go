//--Summary:
//  Create a program that can read text from standard input and count the
//  number of letters present in the input.
//
//--Requirements:
//* Count the total number of letters in any chosen input
//* The input must be supplied from standard input
//* Input analysis must occur per-word, and each word must be analyzed
//  within a goroutine
//* When the program finishes, display the total number of letters counted
//
//--Notes:
//* Use CTRL+D (Mac/Linux) or CTRL+Z (Windows) to signal EOF, if manually
//  entering data
//* Use `cat FILE | go run ./exercise/sync` to analyze a file
//* Use any synchronization techniques to implement the program:
//  - Channels / mutexes / wait groups

package main

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"unicode"
)

type Counter struct {
	count int
	sync.Mutex
}

func parseWord(word string, counter *Counter, wg *sync.WaitGroup) {
	defer wg.Done()

	sum := 0
	for _, letter := range word {
		if unicode.IsLetter(letter) {
			sum++
		}
	}
	counter.Lock()
	defer counter.Unlock()
	fmt.Printf("%s: %d\n", word, sum)
	counter.count += sum

}

func main() {
	dat, err := os.ReadFile("text.txt")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	counter := Counter{}
	var wg sync.WaitGroup

	strData := strings.Split(string(dat), "\n")
	for _, line := range strData {
		words := strings.Split(line, " ")
		for _, word := range words {
			wg.Add(1)
			go parseWord(word, &counter, &wg)
		}
	}

	wg.Wait()

	fmt.Printf("Total letters: %d\n", counter.count)
}
