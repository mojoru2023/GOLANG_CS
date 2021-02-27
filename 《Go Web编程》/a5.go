package main

import "fmt"

func myFunc(){
	i :=0
Here:
	println(i)
	i++
	goto Here
}

func add1(a *int)int  {
	*a = *a+1
	return *a

}

//传指针的好处
// 1. 传指针使得多个函数能操作同一个对象
// 2. 传
func main()  {
	x := 3
	fmt.Println("x=",x)
	x1 := add1(&x)
	fmt.Println("x+1 =",x1)
	fmt.Println("x =",x)
	//var array =[...]byte{'a','b','c','d','e','f','g','h','i','j'}
	//var aSlice,bSlice []byte
	//aSlice = array[:3]
	//aSlice= array[:]
	//bSlice= aSlice[1:3]
	//fmt.Println(bSlice)
	//numbers := make(map[string]int)
	//numbers["one"]=1
	//numbers["ten"]=10
	//fmt.Println(numbers)
	//rating :=map[string]float32 {"c":5,"Go":4,"Python":4}
	//cC,ok := rating["C#"]
	//if ok{
	//	fmt.Println(cC)
	//}else {
	//	fmt.Println("failed~")
	//}
	//del
	// ete(rating,"Go")
	//sum := 0;
	//for index:=0;index < 1000 ;index++  {
	//	sum += index
	//}
	//fmt.Println("sum is equal to",sum)
	//i:=10
	//switch i {
	//case 1:
	//	fmt.Println()
	//}

}