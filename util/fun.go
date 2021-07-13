package util

import (
	"fmt"
	"math"
	"strconv"
)

func Equal(ben float64, n float64, pay float64) float64 {
	var left float64 = 0
	var right float64 = 1
	var mid float64
	var res float64
	limit := 0.01
	for{
		mid = (right + left)/2
		res = Row(ben,n,mid)
		if math.Abs(pay - res) > limit {
			if pay > res {
				left = mid
			}else{
				right = mid
			}
		}else{
			break
		}
	}
	return Decimal(mid*100)
}

func Row (ben float64, n float64, x float64) float64 {
	fen := ben*x*(1+x)*(math.Pow(1+x,n))
	mu := math.Pow(1+x, n-1)
	return fen/mu
}

func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.f", value), 64)
	return value
}