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


	var t2 test=&test2{}
var pt2 test=nil



	fmt.Println(pt2)



	fmt.Println(t2)


}

func (this *test1) add(){

}