package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <msg>\n", os.Args[0])
		os.Exit(1)
	}

	msg := os.Args[1]
	l := utf8.RuneCountInString(msg)

	f := strings.Repeat("!", l)
	s := f + strings.ToUpper(msg) + f

	fmt.Println(s)
}
