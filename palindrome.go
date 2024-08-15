package main

import (
	"fmt"
)

func main() {
	fmt.Println(palindrome("hello"))
	fmt.Println(palindrome("goog"))
}
func palindrome(word string) (b bool) {
	var singleChar []rune = nil
	b = false
	singleChar = []rune(word)

	// go right match all runes to left
	for i := 0; i < len(word)/2; i++ {

		if singleChar[i] == singleChar[len(word)-i-1] {
			fmt.Println("single char i " + string(singleChar[i]))
			fmt.Println("single char len-i " + string(singleChar[len(singleChar)-i-1]))
			b = true
		}

	}

	return b

}
