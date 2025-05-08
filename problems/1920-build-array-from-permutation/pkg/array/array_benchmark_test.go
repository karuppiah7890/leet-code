package array

import "testing"

func BenchmarkBuildArray(b *testing.B) {
	for index := 0; index < b.N; index++ {
		buildArray([]int{6, 11, 9, 4, 7, 10, 0, 8, 5, 3, 2, 17, 12, 13, 14, 15, 1, 16})
	}
}
