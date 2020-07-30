package algosort

var remainingLength int

// bubble sort
func Bubble(sequence []int) []int {
	for i := len(sequence) - 1; i >= remainingLength; i-- {
		if i != 0 {
			if sequence[i-1] > sequence[i] {
				sequence[i-1], sequence[i] = sequence[i], sequence[i-1]
			}
		}
	}
	if len(sequence) != remainingLength {
		remainingLength++
		sequence = Bubble(sequence)
	}
	return sequence
}
