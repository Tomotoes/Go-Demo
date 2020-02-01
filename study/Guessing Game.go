package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	var guess int
	var count int

	num := rand.Intn(100) + 1

	fmt.Println(">> I'm thinking of a number between 1-100 <<")

	for {
		fmt.Print("Guess: ")
		_, err := fmt.Scanf("%d", &guess)
		if err != nil {
			fmt.Println(">> Please input a number")
			continue
		}
		count++
		if guess > num {
			fmt.Println(">> Too high")
		} else if guess < num {
			fmt.Println(" >> Too low")
		} else {
			fmt.Printf("Correct! It took you %d guess!\n", count)
			break
		}

	}
}
