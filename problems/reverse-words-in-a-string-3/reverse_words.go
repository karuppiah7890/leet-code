package reversewords

import (
	"fmt"
	"strings"
)

func ReverseWords(s string) string {
	reversedWords := []string{}
	words := strings.Split(s, " ")

	for _, word := range words {
		reversedWord := reverse(word)
		reversedWords = append(reversedWords, reversedWord)
	}

	return strings.Join(reversedWords, " ")
}

func reverse(s string) string {
	reversedString := strings.Builder{}
	for index := len(s) - 1; index >= 0; index-- {
		c := s[index]
		err := reversedString.WriteByte(c)
		if err != nil {
			panic(fmt.Sprintf("An error occurred while writing byte %v (string: %v) to result string: %v", c, string(c), err))
		}
	}

	return reversedString.String()
}
