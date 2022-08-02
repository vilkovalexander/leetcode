package main

import (
	"fmt"
)

func groupAnagrams(strs []string) [][]string {
	dict := make(map[[26]int][]string, 0)
	result := make([][]string, 0)
	for _, str := range strs {
		tmp := encodeToArray(str)
		if _, ok := dict[tmp]; !ok {
			dict[tmp] = make([]string, 0)
		}
		dict[tmp] = append(dict[tmp], str)
	}
	for _, val := range dict {
		result = append(result, val)
	}
	return result

}

func encodeToArray(str string) [26]int {
	arrayHash := [26]int{}
	for _, char := range str {
		arrayHash[int(char)-'a']++
	}
	return arrayHash
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
