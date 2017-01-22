package main

import (
	"bootstrap"
	channel_ "kernel/intf/external/common"
	"fmt"
	"kernel/channel"
	"kernel/handler"
)

type TestServerChannelHandler struct {
	//impl IChannelInboundHandler
}

func (this *TestServerChannelHandler)ChannelActive_(ctx channel_.IChannelHandlerContext) {
	fmt.Println("server channel active.")
}

func (this *TestServerChannelHandler)ChannelInactive_(ctx channel_.IChannelHandlerContext) {

}

func (this *TestServerChannelHandler)ChannelRead_(ctx channel_.IChannelHandlerContext, msg interface{}) {

}

func (this *TestServerChannelHandler)ExceptionCaught(ctx channel_.IChannelHandlerContext, err error) {
	fmt.Println("ExceptionCaught")
	fmt.Println("has error")
	fmt.Println(err)
}

func main() {

	serverBootstrap := bootstrap.NewServerBootstrap()
	serverBootstrap.Channel(channel.NewTCPServerSocketChannel()).Handler(handler.NewChannelInitializerHandler(func(channel channel_.IChannel) {
		fmt.Println("hello")
	})).ChildHandler(handler.NewChannelInitializerHandler(func(channel channel_.IChannel) {
		channel.Pipeline().AddLast(&TestServerChannelHandler{})
	})).Bind(1024).Sync()

}