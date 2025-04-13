package strings

func LengthOfLongestSubstring(s string) int {
	start, maxLen, seen := 0, 0, make(map[rune]int)
	for i, r := range s {
		if index, ok := seen[r]; ok && index >= start {
			start = index + 1
		}
		seen[r] = i
		if i-start+1 > maxLen {
			maxLen = i - start + 1
		}
	}
	return maxLen
}
