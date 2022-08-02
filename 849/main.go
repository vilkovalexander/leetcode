package main

import (
	"fmt"
)

func maxDistToClosest(seats []int) int {
	i := 0
	j := 0
	maxDistanse := 0

	for ; i < len(seats); i++ {
		if seats[i] == 1 {
			if seats[j] == 0 && maxDistanse < i-j {
				maxDistanse = i - j
			} else if seats[j] != 0 && maxDistanse < (i-j)/2 {
				maxDistanse = (i - j) / 2
			}
			j = i
		}
	}
	if seats[i-1] == 0 {
		if i-j-1 > maxDistanse {
			maxDistanse = i - j - 1
		}
	}
	return maxDistanse
}

func main() {
	fmt.Println(maxDistToClosest([]int{1, 0, 0, 0}))
}
