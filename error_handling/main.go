package main

import (
	"err/custom"
)

func main() {
	custom.PrintByDivisor(10, 2, 5)
	custom.PrintByDivisor(15, 1, 5)
	custom.PrintByDivisor(1, 3, -1)
	custom.PrintByDivisor(140, 5, 0)
}