package main

type Skills []string
type Human struct {
	name string
	age int
	weight int
}

type Student struct {
	Human
	Skills
	int
	spec string
}

