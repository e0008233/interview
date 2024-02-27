package main

import (
	"fmt"
)

func main() {
	temp := make([]int, 4)
	temp2 := append(temp, 4)

	fmt.Println(temp)
	fmt.Println(temp2)

}
