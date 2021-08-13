package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	turns = 5
	limit = 10

	usage = `Welcome to the Lucky Number Game!
The program will pick %d random numbers.
Your mission is to guess one of those numbers.

Usage: %s [guess]
`
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf(usage, limit, os.Args[0])
		return
	}

	guess, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Your guess must be a number.")
		return
	}

	rand.Seed(time.Now().UnixNano())

	offset := limit / 2
	if guess < 0 {
		offset = -offset
	}

	for i := 0; i < turns; i++ {
		n := rand.Intn(limit) + guess - offset
		fmt.Printf("Guess #%d: %d.\n", i+1, n)
		if n == guess {
			fmt.Println("You win!")
			return
		}
	}

	fmt.Println("You lose.")
}
