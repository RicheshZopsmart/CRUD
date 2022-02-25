package hello

import (
	"fmt"

	mapreduce "github.com/RicheshZopsmart/Learning-go/MapReduce"
)

func z() {
	a := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(mapreduce.MyReduce(a, mapreduce.Square))
}
