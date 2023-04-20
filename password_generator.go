package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
	"os"
)

func pwgen(length int, excludeDigits bool, excludeUppercase bool, excludeLowercase bool, excludeSpecialChar bool) string {
	var password string
	chars := ""
	if !excludeDigits {
		chars += "0123456789"
	}
	if !excludeUppercase {
		chars += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	if !excludeLowercase {
		chars += "abcdefghijklmnopqrstuvwxyz"
	}
	if !excludeSpecialChar {
		chars += "!#$%&+*-./"
	}
	if chars == "" {
		chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz%$!&?§#*"
	}

	for i := 0; i < length; i++ {
		randomIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		password += string(chars[randomIndex.Int64()])
	}

	return password
}

func main() {
	length := flag.Int("L", 7, "password length")
	excludeDigits := flag.Bool("d", false, "exclude digits")
	excludeUppercase := flag.Bool("uc", false, "exclude uppercase letters")
	excludeLowercase := flag.Bool("lc", false, "exclude lowercase letters")
	excludeSpecialChar := flag.Bool("sc", false, "exclude special characters")
	flag.Parse()

	password := pwgen(*length, *excludeDigits, *excludeUppercase, *excludeLowercase, *excludeSpecialChar)
	fmt.Println("Generated password:", password)
}

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "Options:")
		flag.PrintDefaults()
	}
}
