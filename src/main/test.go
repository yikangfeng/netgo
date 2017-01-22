package main

import "fmt"

type itest interface {
	add()
}

type test struct {
	itest
}

func (this *test) add() {
	fmt.Println("add method called.")
}

func main() {
	testM()
	var aa string = "hello server"
	fmt.Println([]byte(aa))
	var a itest = &test{}
	fmt.Println(a)
	a.(*test).add()

	s := []string{"a", "b"}
	loopSlice(s...)//one by one transfer.

	arr := [...]int{1, 2, 3, 4, 5}//静态数组arr
	s2 := arr[:]
	fmt.Println(arr[1])
	s2[1] = 22
	fmt.Println(arr[1])

	var pv int = 6
	var pp *int = (&pv)

	fmt.Println(*pp)
}
func loopSlice(s ...string) {
	for _, ss := range s {
		fmt.Println(ss)
	}
}

func testM() {
	defer fmt.Println("testM")
}