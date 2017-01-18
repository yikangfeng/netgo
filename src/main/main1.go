package main

import (

	"fmt"
	channel_ "kernel/intf/external/channel"
	"bootstrap"
	"kernel/channel"
	"kernel/handler"
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

  bootstrap:= bootstrap.NewBootstrap()
	bootstrap.Channel(channel.NewTCPSocketChannel()).Handler(handler.NewChannelInitializerHandler(func(channel channel_.IChannel){
		fmt.Println("hello client")
	})).Connect("127.0.0.1",1024).Sync()

	var a string="b"
	var b string="a"
	if a!=b{
		fmt.Println("true")
	}


}

func (this *test1) add(){

}