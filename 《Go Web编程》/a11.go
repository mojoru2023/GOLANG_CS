package main

import "fmt"


type Rectanle struct {
	width,height float64
}

func area(r Rectanle) float64  {
	return r.height*r.width
}

func main()  {
	r1 := Rectanle{12,3}
	fmt.Println(area(r1))
}