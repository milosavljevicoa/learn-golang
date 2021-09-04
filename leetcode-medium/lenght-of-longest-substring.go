package leetcodemedium

import (
	"fmt"
	"strings"
)

//https://leetcode.com/problems/longest-substring-without-repeating-characters/
func lengthOfLongestSubstring(s string) int {
	maxSubStringLen := 0
	currSubStringLen := 0
	substring := ""

	for i := 0; i < len(s)-1; i++ {
		currSubStringLen = 0
		for j := i + 1; j < len(s); j++ {
			if strings.Contains(substring, string(s[j])) {
				if maxSubStringLen < currSubStringLen {
					maxSubStringLen = currSubStringLen
					break
				}
			} else {
				substring += string(s[j])
				currSubStringLen++
			}
		}
		substring = ""
	}

	return maxSubStringLen
}

func RunlengthOfLongestSubstring() {
	fmt.Println(lengthOfLongestSubstring("abuoueoaujsaeodaeotuhoethkdnj"))
	//runAddTwoNumbers()
}
