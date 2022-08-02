package main

func rotate(matrix [][]int) {
	left := 0
	right := len(matrix) - 1
	for left < right {
		for i := 0; i < right-left; i++ {
			top := left
			bottom := right
			tmp := matrix[top][left+i]

			matrix[top][left+i] = matrix[bottom-i][left]

			matrix[bottom-i][left] = matrix[bottom][right-i]

			matrix[bottom][right-i] = matrix[top+i][right]

			matrix[top+i][right] = tmp
		}
		right -= 1
		left += 1
	}
}

func abs(a, b int) int {
	res := a - b
	if res < 0 {
		return -res
	}
	return res
}

func main() {
	rotate([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
}
