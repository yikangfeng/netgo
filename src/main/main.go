package main

import (
	"fmt"
        "kernel"
	"sync"
	"time"
	"bootstrap"
	"bootstrap/client"
)

func main(){
   var c bootstrap.NetGoContext;
c.ServerName="kf"
	fmt.Println(c.ServerName)

	  const a int=1



	var test string
	test="abc"
	test1:=&test
	t(test1)
	fmt.Println(*test1)

	var m map[string]interface{}=make(map[string]interface{})
	_,ok:=m["a"]
	fmt.Println(ok)

	d:=kernel.GetTCPChannel()
	fmt.Println(d.Name)

	_bootstrap:=bootstrap.New();
	_bootstrap.Channel(kernel.GetTCPChannel()).Bind(1024).Sync()


client.New()


}

func t( b *string)  {
	*b="c"
   fmt.Println(*b)
}

var _lock sync.Mutex
//var _wait sync.WaitGroup
func reentrantLock()  {

     testReentrantlock1()
}

func testReentrantlock(){
	_lock.Lock()

		fmt.Println("lock exec")
		time.Sleep(25)




	_lock.Unlock()
}


func testReentrantlock1(){
	_lock.Lock()
	 testReentrantlock()
		fmt.Println("lock exec1")
		time.Sleep(10*time.Second)




	_lock.Unlock()

}
