package algosort

import "testing"

var sequenceSlice = []int{3, 4, 5, 6, 1, 10, 2, 9, 8, 132, 453, 33232, 22, 55, 99, 65}

func BenchmarkBubble(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bubble(sequenceSlice)
	}
}
