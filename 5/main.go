package main

import "fmt"

func longestPalindrome(s string) string {
	max := 1

	if len(s) == 1 {
		return s
	}
	res := string(s[0])
	for i := range s {
		l, r := i-1, i+1
		count := 0
		for l >= 0 && r < len(s) && s[l] == s[r] {
			count = r - l + 1
			if count > max {
				max = count
				res = s[l : r+1]
			}
			l--
			r++
		}
		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			count = r - l + 1
			if count > max {
				max = count
				res = s[l : r+1]
			}
			l--
			r++
		}
	}
	return res
}

func main() {
	fmt.Println(longestPalindrome("babad"))
}
