package hello

import (
	"fmt"
)

func IsPallindrome(x string) bool {
	if x == "" {
		fmt.Println("Empty String")
		return false
	}
	var start int = 0
	var end int = len(x) - 1
	// fmt.Println(end)
	for start < end {
		if x[start] != x[end] {
			return false
		}
		start++
		end--
	}
	return true
}
func main() {

	var x string
	fmt.Scanln(&x)
	fmt.Println(IsPallindrome(x))

}
