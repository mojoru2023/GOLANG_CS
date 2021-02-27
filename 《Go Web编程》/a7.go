package main

import "os"

var user = os.Getenv("USER")

func init()  {
	if user == ""{
		panic("no value for $USER")
	}
}

func throwPani (f func()) (b bool)  {
	defer func() {
		if x:= recover();x!=nil{
			b =true
		}
	}()
	f()
	return
}


