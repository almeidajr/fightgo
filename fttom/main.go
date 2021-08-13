package main

import (
	"fmt"
	"os"
	"strconv"
)

const usage = `Feet to Meters
--------------
This program converts feet to meters.

Usage: %s <feet>
`

func main() {
	if len(os.Args) < 2 {
		fmt.Printf(usage, os.Args[0])
		return
	}

	feet, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Printf("error: %q is not a number.\n", os.Args[1])
		os.Exit(1)
	}

	meters := feet * 0.3048
	fmt.Printf("%g ft is %g m.\n", feet, meters)
}
