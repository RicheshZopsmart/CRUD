package MapReduce

func MyReduce(a []int, f func(int) int) []int {
	var res []int
	for idx := range a {
		res = append(res, f(a[idx]))
	}
	return res
}
func Square(a int) int {
	return a * a
}
func Cube(a int) int {
	return square(a) * a
}

// func sqrt(a int) float64 {
// 	return (float64(a) / float64(square(a)))
// }

// func main() {
// 	a := []int{1, 2, 3, 4, 5, 6}
// 	fmt.Println(MyReduce(a, square))
// 	fmt.Println(MyReduce(a, cube))
// 	// fmt.Println(MyReduce(a, sqrt))
// }
