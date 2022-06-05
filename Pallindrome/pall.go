package Pallindrome

func isPall(a string) bool {

	var i int = 0
	var j int = len(a) - 1
	if j < 0 {
		return true
	}
	for i < j {
		if a[i] != a[j] {
			return false
		}
		i++
		j--
	}
	return true
}
