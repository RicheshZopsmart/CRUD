package MapReduce

// Accepts Integer List and passes it on to appropriate filters
func Map(integerList []int, f func(interface{}) interface{}) []int {
	var res []int
	for _, val := range integerList {
		res = append(res, f(val))
	}
	return res
}

func evenFilter(int x) int {

}
