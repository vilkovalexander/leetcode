package main

import (
	"fmt"
	"sort"
)

type ListNode struct {
	Val  string
	Next *ListNode
	Prev *ListNode
}

func findItinerary(tickets [][]string) []string {
	hashmap := make(map[string][]string, 0)
	sort.Slice(tickets, func(i, j int) bool {
		if tickets[i][0] < tickets[j][0] {
			return true
		} else if tickets[i][0] > tickets[j][0] {
			return false
		}
		if tickets[i][1] <= tickets[j][1] {
			return true
		} else if tickets[i][1] > tickets[j][1] {
			return false
		}
		return true

	})
	for _, path := range tickets {
		if _, ok := hashmap[path[0]]; !ok {
			hashmap[path[0]] = make([]string, 0)
		}
		hashmap[path[0]] = append(hashmap[path[0]], path[1])
	}
	//for key, v := range hashmap {
	//	fmt.Println(key, "     ", v)
	//}
	res := make([]string, 1)
	res[0] = "JFK"
	dfs(len(tickets)+1, "JFK", &res, &hashmap)
	return res
}

func dfs(n int, v string, res *[]string, adj *map[string][]string) bool {
	if len(*res) == n {
		return true
	}
	if val, ok := (*adj)[v]; !ok || len(val) == 0 {
		return false
	}
	tmp := (*adj)[v]
	for i, vertex := range tmp {
		if i+1 >= len(tmp) {
			(*adj)[v] = (*adj)[v][:i]
		} else if i == 0 {
			(*adj)[v] = (*adj)[v][1:]
		} else {
			(*adj)[v] = append((*adj)[v][:i], (*adj)[v][i+1:]...)
		}
		*res = append(*res, vertex)
		if dfs(n, vertex, res, adj) {
			return true
		}
		temp := (*adj)[v]
		if i+1 >= len(tmp) {
			(*adj)[v] = append((*adj)[v], vertex)
		} else if i == 0 {
			(*adj)[v] = append([]string{vertex}, temp...)
		} else {
			(*adj)[v] = append((*adj)[v][:i+1], vertex)
			(*adj)[v] = append((*adj)[v], temp[i+1:]...)
		}
		*res = (*res)[:len(*res)-1]
	}
	return false
}

func main() {

	fmt.Println(findItinerary([][]string{{"CBR", "JFK"}, {"TIA", "EZE"}, {"AUA", "TIA"}, {"JFK", "EZE"}, {"BNE", "CBR"}, {"JFK", "CBR"}, {"CBR", "AUA"}, {"EZE", "HBA"}, {"AXA", "ANU"}, {"BNE", "EZE"}, {"AXA", "EZE"}, {"AUA", "ADL"}, {"OOL", "JFK"}, {"BNE", "AXA"}, {"OOL", "EZE"}, {"EZE", "ADL"}, {"TIA", "BNE"}, {"EZE", "TIA"}, {"JFK", "AUA"}, {"AUA", "EZE"}, {"ANU", "ADL"}, {"TIA", "BNE"}, {"EZE", "OOL"}, {"ANU", "BNE"}, {"EZE", "ANU"}, {"ANU", "AUA"}, {"BNE", "ANU"}, {"CNS", "JFK"}, {"TIA", "ADL"}, {"ADL", "AXA"}, {"JFK", "OOL"}, {"AUA", "ADL"}, {"ADL", "TIA"}, {"ADL", "ANU"}, {"ADL", "JFK"}, {"BNE", "EZE"}, {"ANU", "BNE"}, {"JFK", "BNE"}, {"EZE", "AUA"}, {"EZE", "AXA"}, {"AUA", "TIA"}, {"ADL", "CNS"}, {"AXA", "AUA"}}))
}
