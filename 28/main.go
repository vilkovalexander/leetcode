package main

import "fmt"

func strStr(haystack string, needle string) int {
	i := 0
	j := len(needle) - 1
	for j < len(haystack) {
		m := 0
		for k := i; k <= j; k++ {
			if haystack[k] != needle[m] {
				j++
				i++
				break
			}
			m++
		}
		if m == len(needle) && j < len(haystack) && needle[m-1] == haystack[j] {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(strStr("mississippi",
		"issip"))
}
