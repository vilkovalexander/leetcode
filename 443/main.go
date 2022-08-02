package main

import (
	"fmt"
	"strconv"
	"strings"
)

func compress(chars []byte) int {
	count := 1
	realPos := 0
	comPos := 0
	for i := 1; i < len(chars); i++ {
		if chars[i] == chars[comPos] {
			count++
		} else {
			chars[realPos] = chars[comPos]
			realPos++
			if count != 1 {
				for _, val := range countToSliceStrings(count) {
					chars[realPos] = []byte(val)[0]
					realPos++
				}
			}
			comPos = i
			count = 1
		}
	}
	return realPos
}

func countToSliceStrings(count int) []string {
	str := strconv.Itoa(count)
	return strings.Split(str, "")
}

func printdd(chars []byte) {
	for _, val := range chars {
		fmt.Printf("%c ", val)
	}
	fmt.Println()
}

func main() {
	chars := []byte{'a', 'b', 'c', 'c', 'c', 'c', 'c', 'c'}
	chars = append(chars, '~')
	printdd(chars)
	compress(chars)
	printdd(chars)
}
