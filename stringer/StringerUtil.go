package main

import (
	"fmt"
)

type Rectangle struct {
	Length int
	width  int
}

func (r Rectangle) String() string {
	return fmt.Sprintf("Length = %v, Width = %v", r.Length, r.width)
}

func main() {
	r := Rectangle{
		Length: 10,
		width:  10,
	}
	fmt.Println(r)
}
