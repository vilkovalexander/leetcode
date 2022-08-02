package main

import "fmt"

func generateParenthesis(n int) []string {
	stack := make([]byte, 0)
	result := make([]string, 0)
	backtrack(0, 0, n, stack, &result)
	return result
}

func backtrack(opened, closed, n int, stack []byte, result *[]string) {
	if opened == closed && closed == n {
		*result = append((*result), string(stack))
		return
	}
	if opened < n {
		stack = append(stack, '(')
		backtrack(opened+1, closed, n, stack, result)
		stack = stack[:len(stack)-1]
	}
	if closed < opened {
		stack = append(stack, ')')
		backtrack(opened, closed+1, n, stack, result)
		stack = stack[:len(stack)-1]
	}
}
func main() {
	fmt.Println(generateParenthesis(3))
}
