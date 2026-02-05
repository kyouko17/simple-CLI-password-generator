package password

import (
	"crypto/rand"
	"errors"
	"math/big"
)

const (
	whitelist     string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@$#%/?-~"
	lowerCase     string = "abcdefghijklmnopqrstuvwxyz"
	upperCase     string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers       string = "1234567890"
	uniqueSymbols string = "!@$#%/?-~"
)

var ErrNotAcceptableLength = errors.New("password length must be between 4 and 32")

func Generate(length int) (string, error) {
	//Error handle
	if length < 4 || length > 32 {
		return "", ErrNotAcceptableLength
	}

	var res []byte

	//Add base for pass
	err := addRandomCharFromString(&res, lowerCase)
	if err != nil {
		return "", err
	}
	err = addRandomCharFromString(&res, upperCase)
	if err != nil {
		return "", err
	}
	err = addRandomCharFromString(&res, numbers)
	if err != nil {
		return "", err
	}
	err = addRandomCharFromString(&res, uniqueSymbols)
	if err != nil {
		return "", err
	}

	//Fill remaining
	for i := 0; i < length-4; i++ {
		err := addRandomCharFromString(&res, whitelist)
		if err != nil {
			return "", err
		}
	}

	//Shuffle (Fisher-Yates)
	for i := len(res) - 1; i > 0; i-- {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
		if err != nil {
			return "", err
		}
		res[i], res[n.Int64()] = res[n.Int64()], res[i]
	}

	return string(res), nil
}

func addRandomCharFromString(res *[]byte, source string) error {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(len(source))))
	if err != nil {
		return err
	}
	*res = append(*res, source[n.Int64()])
	return nil
}
