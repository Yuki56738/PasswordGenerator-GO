package main

import (
	"crypto/rand"
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("Hello, world!")
	const SEED = "0123456789abcdefghijklmnopqrstuvwxyz"
	SeedLen := int64(utf8.RuneCountInString(SEED))

	for i := 0; i < 8; i++ {
		n, err := rand.Int(rand.Reader(), big.newInt(SeedLen))
		if err != nil {
			return fmt.Errorf(":%w", err)
		}
	}

}
