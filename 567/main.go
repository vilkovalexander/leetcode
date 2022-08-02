package main

import (
	"fmt"
)

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	containString := [26]int{}
	for _, s := range s1 {
		containString[s-'a'] += 1
	}
	diff := [26]int{}
	for i := 0; i < len(s1); i++ {
		diff[s2[i]-'a'] += 1
	}
	if containString == diff {
		return true
	}
	for i := 1; i <= len(s2)-len(s1); i++ {
		diff[s2[i-1]-'a']--
		diff[s2[i+len(s1)-1]-'a']++
		if containString == diff {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(checkInclusion("adc", "dcda"))

}
