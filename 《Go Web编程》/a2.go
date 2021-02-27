package main

import "fmt"

func main(){
	//str------[]byte-------str
	var s string = "hello"
	c := []byte(s)
	c[0] = 'c'
	s2 := string(c)
	fmt.Println(s2)
	fmt.Printf("%s \n",s2)

}
