package main

import "fmt"

type Human struct {
	name string
	age int
	weight int
}

type Student struct {
	Human
	speciality string
}

func main()  {
	mark := Student{Human{"Mark",25,220},"COA"}
	fmt.Println(mark.age,mark.name)
}