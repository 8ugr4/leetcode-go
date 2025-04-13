package problems

import (
	"LeetCodeGo/problems/strings"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoSums(t *testing.T) {
	result := TwoSum([]int{10, 20, 30, 40}, 30)
	assert.Equal(t, []int{0, 1}, result)
}

func Test_lengthOfLongestSubstring(t *testing.T) {
	result := strings.LengthOfLongestSubstring("abcabcbb")
	assert.Equal(t, 3, result)
	result = strings.LengthOfLongestSubstring("bbbbb")
	assert.Equal(t, 1, result)
	result = strings.LengthOfLongestSubstring("pwwkew")
	assert.Equal(t, 3, result)
	result = strings.LengthOfLongestSubstring("abcb")
	assert.Equal(t, 3, result)
}
