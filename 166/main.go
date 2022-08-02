package main

import "strconv"

func fractionToDecimal(numerator int, denominator int) string {
	c := float64(numerator) / float64(denominator)
	st := strconv.FormatFloat()
}
