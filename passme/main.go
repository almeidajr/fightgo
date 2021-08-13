package main

import (
	"fmt"
	"os"
)

const (
	usage    = "Usage: %s [username] [password]\n"
	errUser  = "Access denied for %q.\n"
	errPwd   = "Invalid password for %q.\n"
	accessOk = "Access granted to %q.\n"

	user = "admin"
	pass = "0000"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf(usage, os.Args[0])
		return
	}

	u, p := os.Args[1], os.Args[2]
	if u != user {
		fmt.Printf(errUser, u)
	} else if p != pass {
		fmt.Printf(errPwd, u)
	} else {
		fmt.Printf(accessOk, u)
	}
}
