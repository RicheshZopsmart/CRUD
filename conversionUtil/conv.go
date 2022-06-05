package main

import (
	"fmt"
	"strconv"
)

func main() {
	i, err := strconv.Atoi("-44")
	if err != nil {
		panic(err)
	}
	fmt.Println(i)
	s := strconv.Itoa(i)
	fmt.Println(s)
}
