package main

import (

	// "github.com/RicheshZopsmart/Learning-go/Sort"

	"fmt"

	in "github.com/RicheshZopsmart/Learning-go/interfaces_imp"

	errpack "github.com/RicheshZopsmart/Learning-go/ErrorExample"

	readers "github.com/RicheshZopsmart/Learning-go/ReadersUtil"
)

// func main() {
// 	var a = []int{1, 2, 3, -1, -77, 44, 33, 22}
// 	fmt.Println(Sort.BubbleSort(a))
// 	var low int = 0
// 	var high int = len(a) - 1
// 	Sort.MergeSort(a, low, high)
// 	fmt.Println(a)
// }

func main() {

	var s in.StudentIterface
	s = in.Student{Name: "Richesh", Age: 21, Marks: 0}
	fmt.Println(s.GetAge())
	fmt.Println(s.IsPass())
	s.PrintDetails()

	fmt.Println(errpack.Add(2.0, 3))
	fmt.Println(errpack.Divide(5, 2))
	fmt.Println(errpack.Divide(5, 0))

	x := readers.StringStream()
	if x != nil {
		fmt.Println(x)
	}
}
