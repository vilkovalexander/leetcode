package main

import "fmt"

func numTeams(rating []int) int {
	if len(rating) < 3 {
		return 0
	}
	num := make([]int, len(rating))
	for l, _ := range rating {
		num[len(rating)-l-1] = rating[l]
	}
	count := 0
	count += fdsfds(rating)
	count += fdsfds(num)
	return count
}

func fdsfds(rating []int) int {
	i := 0
	j := 1
	k := 2
	res := 0
	for k < len(rating) && j < len(rating) && i < len(rating) {
		if rating[i] > rating[j] {
			j++
			k++
			continue
		}
		for c := k; c < len(rating); c++ {
			if rating[c] > rating[j] {
				res++
			}
		}
		i++
		j++
		k++
	}
	return res
}

func main() {
	fmt.Println(numTeams([]int{1, 2, 3, 4}))
}
