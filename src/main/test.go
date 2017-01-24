package main

import (
	"fmt"
	"encoding/gob"
	"bytes"
	"encoding/binary"

)

type itest interface {
	add()
}

type test struct {
	itest
}

func (this *test) add() {
	fmt.Println("add method called.")
}

type testoutstruct struct {
	Name string
}

func testout(out interface{}) {

	var a string="kf"
	strp:=out.(*string)
	fmt.Println(strp)
	strp=&a
	fmt.Println(strp)
	fmt.Println(*strp)
}

func testslice(out []string) {
	t := append(out, "aa")
	fmt.Println(t)
}
func main() {

	var ss string="test"

	testouts:=&ss

	fmt.Println(testouts)

	testout(testouts)
        fmt.Println(*testouts)




	var network1 bytes.Buffer
	packetBody := "hello server"
	//var packetLength bytes.Buffer
	binary.Write(&network1, binary.BigEndian, uint32(len(packetBody)))

	binary.Write(&network1, binary.BigEndian, []byte("<BOF>"))

	binary.Write(&network1, binary.BigEndian, []byte("stringprotocol"))

	binary.Write(&network1, binary.BigEndian, []byte(packetBody))

	var bb uint32
	bb = binary.BigEndian.Uint32(network1.Bytes()[0:4])
	fmt.Println(bb)

	fmt.Println(network1.Len())
	fmt.Println(network1.Len() - 4 - 5 - int(bb))

	fmt.Println(string(network1.Bytes()[9:9 + 14]))

	fmt.Println("end")

	slice := make([]string, 0)
	slice1 := append(slice, "a", "b", "c", "d")
	fmt.Println(slice1)

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, []byte("<BOF>"))
	fmt.Println("hold byte size=")
	fmt.Println(buf.Len())

	var network bytes.Buffer
	enc := gob.NewEncoder(&network)
	enc.Encode("hello server")

	dec := gob.NewDecoder(&network)

	var t string;
	dec.Decode(&t)
	fmt.Println("t=" + t)

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