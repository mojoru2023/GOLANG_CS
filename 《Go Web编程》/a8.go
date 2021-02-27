package main

import "fmt"

type person struct {
	name string
	age int
}

func Older(p1,p2 person)(person,int)  {
	if p1.age>p2.age{
		return p1,p1.age-p2.age
	}

}


func main()  {
	//var P person
	//P.name ="GREEE"
	//P.age=66
	//p = person{name:"GAS",age:66}
	//fmt.Printf("The %s",P.name)

	var tom person
	tom.name,tom.age = "Tom",18
	bob := person{"Bob",66}
	t1,t2 := Older(tom,bob)
	fmt.Printf("%s %s %s",t1,t2)
}