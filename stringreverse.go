package main

import (
	"fmt"
	"math"
)

func FirstReverse(str string) string {
	rns := []rune(str)
	r := len(rns) - 1

	for i := 0; i <= int(math.Ceil(float64(r/2))); i++ {
		x := rns[i]
		rns[i] = rns[r-i]
		rns[r-i] = x
	}

	return string(rns)
}

func main() {
	// do not modify below here, readline is our function
	// that properly reads in the input for you
	fmt.Println(FirstReverse("coderbytes"))
}
