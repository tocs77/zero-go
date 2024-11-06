//--Summary:
//  Create a program that can create a report of rune information from
//  lines of text.
//
//--Requirements:
//* Create a single function to iterate over each line of text that is
//  provided in main().
//  - The function must return nothing and must execute a closure
//* Using closures, determine the following information about the text and
//  print a report to the terminal:
//  - Number of letters
//  - Number of digits
//  - Number of spaces
//  - Number of punctuation marks
//
//--Notes:
//* The `unicode` stdlib package provides functionality for rune classification

package main

import (
	"fmt"
	"unicode"
)

func reportLines(lines []string, fn func(line string) (int, int, int, int)) {
	var letterCount, digitCount, spaceCount, punctuationCount int
	for _, line := range lines {
		lineLetterCount, lineDigitCount, lineSpaceCount, linePunctuationCount := fn(line)
		letterCount += lineLetterCount
		digitCount += lineDigitCount
		spaceCount += lineSpaceCount
		punctuationCount += linePunctuationCount
	}
	fmt.Printf("There are %d letters, %d digits, %d spaces, and %d punctuation marks in these lines of text!\n", letterCount, digitCount, spaceCount, punctuationCount)
}

func countLetters(line string) int {
	count := 0
	for _, r := range line {
		if unicode.IsLetter(r) {
			count++
		}
	}
	return count
}

func countDigits(line string) int {
	count := 0
	for _, r := range line {
		if unicode.IsDigit(r) {
			count++
		}
	}
	return count
}

func countSpaces(line string) int {
	count := 0
	for _, r := range line {
		if unicode.IsSpace(r) {
			count++
		}
	}
	return count
}

func countPunctuation(line string) int {
	count := 0
	for _, r := range line {
		if unicode.IsPunct(r) {
			count++
		}
	}
	return count
}

func analyzeLine(line string) (int, int, int, int) {
	letterCount := countLetters(line)
	digitCount := countDigits(line)
	spaceCount := countSpaces(line)
	punctuationCount := countPunctuation(line)
	return letterCount, digitCount, spaceCount, punctuationCount
}

func main() {
	lines := []string{
		"There are",
		"68 letters,",
		"five digits,",
		"12 spaces,",
		"and 4 punctuation marks in these lines of text!",
	}

	reportLines(lines, analyzeLine)
}
