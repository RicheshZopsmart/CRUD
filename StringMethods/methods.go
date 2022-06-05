package StringMethods

import (
	"fmt"
	"strings"
)

func Execute() {
	s := "RicHesH GuPtA"
	fmt.Println(strings.ToLower(s))
	fmt.Println(strings.ToUpper(s))
	fmt.Println(strings.Split(s, " ")[0])

	j := []string{"Hi"}
	fmt.Println(strings.Join(j, "--"))
	//fmt.Println(j)
	aa := strings.Repeat("A", 22)
	fmt.Println(aa)

}
