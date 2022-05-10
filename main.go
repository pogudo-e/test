package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(MyString.IsUpperCase("hello I AM DONALD")) //ff
	fmt.Println(MyString.IsUpperCase("HELLO I AM DONALD")) // tt

	// fmt.Println(IsPalindrome("abba"))
	// fmt.Println(IsPalindrome("hello"))

}

type MyString string

func (s MyString) IsUpperCase() bool {
	a := []rune(s)
	if ranj(a) > 0 {
		fmt.Println(ranj(a))
		return false
	} else {
		fmt.Println(ranj(a))

		return true
	}
}

func ranj(a []rune) int {
	var b int
	for _, i := range a {
		if !unicode.IsUpper(i) {
			if i != 040 {
				b++
			}
		}
	}
	return b
}
