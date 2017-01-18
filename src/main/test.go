package main

import "fmt"

type itest interface {
	add()
}

type test struct {
	itest
}

func (this *test) add()  {
	fmt.Println("add method called.")
}


func main()  {
	var a itest=&test{}
	fmt.Println(a)
	a.(*test).add()
}