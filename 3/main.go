package main

import (
	"fmt"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	hashmap := make(map[string]struct{}, 0)
	j := 0
	result := 0
	args := strings.Split(s, "")
	for _, symbol := range args {
		if _, ok := hashmap[symbol]; ok {
			for {
				_, ok := hashmap[symbol]
				if ok {
					delete(hashmap, args[j])
					j++
				} else {
					break
				}
			}
			hashmap[symbol] = struct{}{}
		} else {
			hashmap[symbol] = struct{}{}
			result = max(result, len(hashmap))
		}
	}
	result = max(result, len(hashmap))
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(lengthOfLongestSubstring("qrsvbspk"))
}
