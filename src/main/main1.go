package main

import (

	"fmt"
	channel_ "kernel/intf/external/common"
	"bootstrap"
	"kernel/channel"
	"kernel/handler"
	handler_ "kernel/intf/external/handler"
	"sync"
)


type test interface {

	add()
}

type test1 struct {//impl test

}



type test2 struct {
	test1
}

type TestClientChannelHandler struct {
	//impl IChannelInboundHandler
	handler_.IChannelInboundHandler
}

var _ctx handler_.IChannelHandlerContext

func (this *TestClientChannelHandler)ChannelActive_(ctx handler_.IChannelHandlerContext) {
	fmt.Println("TestClientChannelHandler channel active called.")
	_ctx=ctx
	_wait.Done()
}

func (this *TestClientChannelHandler)ChannelInactive_(ctx handler_.IChannelHandlerContext) {

}

func (this *TestClientChannelHandler)ChannelRead_(ctx handler_.IChannelHandlerContext, msg interface{}) {

}

var _wait sync.WaitGroup

func main(){
  bootstrap:= bootstrap.NewBootstrap()
	bootstrap.Channel(channel.NewTCPSocketChannel()).Handler(handler.NewChannelInitializerHandler(func(channel channel_.IChannel){
		channel.Pipeline().AddLast(&TestClientChannelHandler{})
	})).Connect("127.0.0.1",1024)

	_wait.Add(1)
	_wait.Wait()



	_ctx.Write([]byte("hello server"))






	var a string="b"
	var b string="a"
	if a!=b{
		fmt.Println("true")
	}


}

func (this *test1) add(){

}