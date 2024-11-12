package main

import (
	"fmt"

	"golang.org/x/exp/rand"
)

func generateRandomInt(min, max int) <-chan int {
	out := make(chan int, 3)
	go func() {
		for {
			out <- rand.Intn(max-min) + min
		}
	}()
	return out
}
func genereateRandomIntN(count, min, max int) <-chan int {
	out := make(chan int, 1)
	go func() {
		for i := 0; i < count; i++ {
			out <- rand.Intn(max-min) + min
		}
		close(out)
	}()
	return out
}

func main() {
	randomInts := generateRandomInt(0, 100)
	fmt.Println("Generate random integers infinitely")
	fmt.Println(<-randomInts)
	fmt.Println(<-randomInts)
	fmt.Println(<-randomInts)

	fmt.Println("Generate 5 random integers")
	randomInts = genereateRandomIntN(5, 0, 100)
	for i := 0; i < 5; i++ {
		fmt.Println(<-randomInts)
	}

	fmt.Println("Generate 5 random integers, check close")
	randomInts = genereateRandomIntN(5, 0, 100)
	for {
		n, ok := <-randomInts
		if !ok {
			break
		}
		fmt.Println(n)
	}
}
