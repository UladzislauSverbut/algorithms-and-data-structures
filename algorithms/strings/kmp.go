package strings

// kmp implements the Knuth-Morris-Pratt algorithm for substring search. It returns the starting index of the first occurrence of the substring in the text, or -1 if the substring is not found.
func kmp(text string, substring string) int {
	pi := buildPiFunction(substring)

	textPos := 0
	substringPos := 0

	for textPos < len(text) && substringPos < len(substring) {
		if text[textPos] == substring[substringPos] {
			textPos++
			substringPos++
		} else {
			if substringPos > 0 {
				substringPos = pi[substringPos-1]
			} else {
				textPos++
			}
		}
	}
	if substringPos == len(substring) {
		return textPos - substringPos
	}

	return -1
}

func buildPiFunction(substring string) []int {
	pi := make([]int, len(substring))

	left := 0
	right := 1

	for right < len(substring) {
		if substring[left] == substring[right] {
			pi[right] = left + 1
			left++
			right++
		} else {
			if left == 0 {
				pi[right] = 0
				right++
			} else {
				left = pi[left-1]
			}
		}
	}

	return pi
}
