package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kyouko17/password_generator/internal/password"
)

func main() {
	length := flag.Int("length", 8, "length of password (4-32)")
	flag.Parse()

	password, err := password.Generate(*length)
	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}
	fmt.Println("generated password:", password)
}
