package main

import (
	"fmt"
	"math/rand"
)

const (
	upperBound   = 10
	linesPerPage = 20
)

func main() {
	for i := 0; i < linesPerPage; i++ {
		fmt.Printf("%d * %d = \n", rand.Intn(upperBound+1), rand.Intn(upperBound+1))
	}
}
