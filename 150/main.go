package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	for _, val := range tokens {
		if val != "*" && val != "/" && val != "+" && val != "-" {
			number, _ := strconv.Atoi(val)
			stack = append(stack, number)
		} else {
			a, b := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			if val == "*" {
				stack = append(stack, a*b)
				continue
			}
			if val == "-" {
				stack = append(stack, a-b)
				continue
			}
			if val == "/" {
				stack = append(stack, a/b)
				continue
			}
			if val == "+" {
				stack = append(stack, a+b)
				continue
			}
		}
	}
	return stack[0]
}

func main() {
	fmt.Println(evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))

}
