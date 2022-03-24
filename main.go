package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func main() {
	input := "The quick brown fox jumped over the lazy dog"
	rev, err := Reverse(input)
	doubleRev, doubleErr := Reverse(rev)
	fmt.Printf("original: %q\n", input)
	fmt.Printf("reversed: %q\n, err: %v\n", rev, err)
	fmt.Printf("reversed back: %q\n, err: %v\n", doubleRev, doubleErr)
}

func Reverse(s string) (string, error) {
	if !utf8.ValidString(s) {
		return s, errors.New("input is not valid UTF-8")
	}
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r), nil
}
