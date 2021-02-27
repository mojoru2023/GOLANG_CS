package main

import "fmt"

func main() {
	arr := [3]int{2, 3}
	arr2 := [12]int{2, 3}
	c := [...]int{4, 5}
	fmt.Println(arr, arr2, c)

	ss := arr2[0:6]
	fmt.Println(ss)


}