package main

import (
	"bootstrap"
	channel_ "kernel/intf/external/common"
	"fmt"
	"kernel/channel"
	"kernel/handler"
	"kernel/handler/transport/protocols"
	string_ "kernel/handler/transport/protocols/string"
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
	fmt.Println("TestServerChannelHandler accept content=" + msg.(string))
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
		channel.Pipeline().AddLast(protocols.NewLengthFieldBasedFrameDecoder(0,0,4,0,0))
		channel.Pipeline().AddLast(string_.NewStringUnpacket())
		channel.Pipeline().AddLast(&TestServerChannelHandler{})
	})).Bind(1024).Sync()

}