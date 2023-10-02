package reversewords

func ReverseWords(s string) string {
	stringBytes := []byte(s)

	wordStart := 0
	wordEnd := 0
	for index := 1; index < len(stringBytes); index++ {

		if stringBytes[index] == ' ' {
			wordEnd = index - 1
			reverse(stringBytes, wordStart, wordEnd)
			wordStart = index + 1
			wordEnd = wordStart
		}

		if index == len(stringBytes)-1 {
			wordEnd = index
			reverse(stringBytes, wordStart, wordEnd)
		}
	}

	return string(stringBytes)
}

// reverse the bytes in a byte array between the fromIndex and toIndex, inclusive
func reverse(s []byte, startIndex int, endIndex int) {
	length := endIndex - startIndex + 1
	midIndex := startIndex + ((length - 1) / 2)
	leftIndex := startIndex
	var value byte

	for rightIndex := endIndex; rightIndex > midIndex; rightIndex-- {
		value = s[leftIndex]
		s[leftIndex] = s[rightIndex]
		s[rightIndex] = value
		leftIndex++
	}
}
