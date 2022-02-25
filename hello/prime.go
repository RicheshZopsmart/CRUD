package hello

import (
	"fmt"
	"math"
)

func isPrime(x int) bool {

	if x == 2 {
		return true
	}

	if x%2 == 0 {
		return false
	}

	lim := math.Sqrt(float64(x))

	for i := 3; i <= int(lim); i += 2 {
		if x%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println("Enter String")
	var s string
	fmt.Scanln(&s)
	fmt.Println(isPrime(s))
}
