package main

import (

	"fmt"


)


type test interface {

	add()
}

type test1 struct {//impl test

}



type test2 struct {
	test1
}

func main(){


	var a string="b"
	var b string="a"
	if a!=b{
		fmt.Println("true")
	}


}

func (this *test1) add(){

}