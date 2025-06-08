package custom

import (
	"fmt"
)
type ZeroDivisionError struct {
	dividend float64
}

func (z ZeroDivisionError) Error() string {
	return fmt.Sprintf("cannot divide %.2f by zero", z.dividend)
}

func Divide(n float64, d float64) (float64, error) {
	if d == 0 {
		return 0, ZeroDivisionError{dividend: n}
	}
	return n / d, nil
}

func PrintByDivisor(x int,  y int, limit int) {
	for limit <= 0 {
		f, e := Divide(float64(x), float64(y))
		if e != nil {
			switch err := e.(type) {
			case ZeroDivisionError:
				fmt.Println("ZeroDivisionError:", err)
			default:
				fmt.Println("Unknown error:", err)
			}
			return
		}
		fmt.Println("Result:", f)
		limit--
		x *= limit
	}
}