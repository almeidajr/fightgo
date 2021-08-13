package main

import (
	"fmt"
	"time"
)

func main() {
	switch h := time.Now().Hour(); {
	case h < 6:
		fmt.Println("good night")
	case h < 12:
		fmt.Println("good morning")
	case h < 18:
		fmt.Println("good afternoon")
	default:
		fmt.Println("good evening")
	}
}
