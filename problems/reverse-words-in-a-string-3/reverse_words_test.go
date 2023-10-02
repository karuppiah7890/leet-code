package reversewords_test

import (
	"testing"

	reversewords "github.com/karuppiah7890/leet-code/problems/reverse-words-in-a-string-3"
	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	input          string
	expectedOutput string
}

type TestCases []TestCase

func TestReverseWords(t *testing.T) {
	testCases := TestCases{
		{
			input:          "Awesome",
			expectedOutput: "emosewA",
		},
		{
			input:          "Awesome Stuff",
			expectedOutput: "emosewA ffutS",
		},
	}

	for _, testCase := range testCases {
		output := reversewords.ReverseWords(testCase.input)

		assert.Equal(t, testCase.expectedOutput, output, "Input: %v", testCase.input)
	}
}

func BenchmarkReverseWords(b *testing.B) {
	for i := 0; i < b.N ; i++ {
		reversewords.ReverseWords("It's Awesome Stuff Over Here")
	}
}
