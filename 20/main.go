package main

import (
	"fmt"
	"strings"
)

func isValid(s string) bool {
	stack := make([]string, 0)
	var ok bool
	var val string
	closeToOpen := map[string]string{")": "(", "}": "{", "]": "["}
	for _, braket := range strings.Split(s, "") {
		if val, ok = closeToOpen[braket]; !ok {

			stack = append(stack, braket)
		} else {
			if stack[len(stack)-1] != val {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}
	if len(stack)%2 == 1 {
		return false
	}
	return true
}

func main() {
	fmt.Println(isValid("()[]{}"))
}
