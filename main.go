package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

var Whitelist string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@$#%/?-~"

func main() {
	fmt.Println(PasswordGenerator(8))
}
func PasswordGenerator(length int) string {
	var sb strings.Builder
	max := big.NewInt(int64(len(Whitelist)))

	for range length {
		n, _ := rand.Int(rand.Reader, max)
		sb.WriteRune(rune(Whitelist[n.Int64()]))
	}

	return sb.String()
}
