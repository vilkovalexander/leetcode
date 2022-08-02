package main

import "fmt"

func isSubsequence(s string, t string) bool {
	if len(s) > len(t) {
		return false
	}
	if len(s) == len(t) {
		return s == t
	}
	j := 0

	for i := 0; i < len(t) && j < len(s); i++ {
		if s[j] == t[i] {
			j++
		}
	}
	return len(s) == j
}

func numMatchingSubseq(s string, words []string) int {
	count := 0
	for _, str := range words {
		if isSubsequence(str, s) {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(isSubsequence("axc", "ahbgdc"))
}
