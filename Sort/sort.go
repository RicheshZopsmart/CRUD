package Sort

func BubbleSort(a []int) []int {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				var temp int = a[i]
				a[i] = a[j]
				a[j] = temp
			}
		}
	}
	return a
}

func MergeSort(a []int, low int, high int) {

	if low < high {
		var mid int = (low + high) / 2
		MergeSort(a, low, mid)
		MergeSort(a, mid+1, high)
		Merge(a, low, mid, high)
	}
}
func Merge(a []int, low int, mid int, high int) {
	temp := make([]int, high-low+1)

	var i int = low
	var j int = mid + 1
	var k int = 0
	for i <= mid && j <= high {

		if a[i] < a[j] {
			temp[k] = a[i]
			i++
		} else {
			temp[k] = a[j]
			j++
		}
		k++

	}

	for i <= mid {
		temp[k] = a[i]
		i++
		k++
	}

	for j <= high {

		temp[k] = a[j]
		j++
		k++
	}

	for i := low; i <= high; i++ {
		a[i] = temp[i-low]
	}

}

// func main() {

// 	var a = []int{1, 2, 3, -1, -77, 44, 33, 22}
// 	// fmt.Println(BubbleSort(a))
// 	var low int = 0
// 	var high int = len(a) - 1
// 	MergeSort(a, low, high)
// 	fmt.Println(a)
// }
