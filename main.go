package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

func main() {
	fmt.Println(CreatePass(16))
}
func CreatePass(len int) string {
	var sb strings.Builder

	for range len {
		rn := randomASCINumber()
		for rn < 33 {
			rn = randomASCINumber()
		}
		sb.WriteString(string(rn))
	}
	return sb.String()
}
func randomASCINumber() rune {
	max := big.NewInt(127)
	n, err := rand.Int(rand.Reader, max)

	if err != nil {
		panic(err)
	}
	return rune(n.Int64())
}
