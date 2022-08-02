package main

import "fmt"

func isPalindrome(s string) bool {
	var i, j int
	res := make([]byte, 0)
	for _, val := range []byte(s) {
		if val >= 65 && val < 91 {
			res = append(res, val+32)
		}
		if (val >= 97 && val <= 122) || (val >= 48 && val <= 57) {
			res = append(res, val)
		}
	}
	if len(res)%2 == 1 {
		i = len(res)/2 - 1
		j = i + 2
	} else {
		i = len(res)/2 - 1
		j = i + 1
	}
	fmt.Println(string(res))
	for i >= 0 && j < len(res) {
		if res[i] != res[j] {
			return false
		}
		i--
		j++
	}
	return true
}

func main() {
	fmt.Println(isPalindrome("ab"))
}

// amanaplanacanalpanama
// amanaplanacanalpanama
