package main

import "fmt"

type testInt func(int) bool

func isOdd(integer int) bool {
	if integer %2 ==0{
		return false
	}
	return true
}

func isEven(integer int)bool{
	if integer %2==0{
		return true
	}
	return false

}

func filter(slice []int,f testInt) []int  {
	var result []int
	for _, value := range slice {
		if f(value){
			result = append(result,value)
		}

	}
	return result
}
func main()  {
	//for i:=0;i<5 ;i++  {
	//	defer fmt.Println(i)
	//}
	slice := []int{1,2,3,4,6}
	fmt.Println(slice)
	odd := filter(slice,isOdd)
	fmt.Println(odd)
	even := filter(slice,isEven)
	fmt.Println(even)
}