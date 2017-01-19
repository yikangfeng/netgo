package main

import (
	"bootstrap"
	channel_ "kernel/intf/external/common"
	"fmt"
	"kernel/channel"
	"kernel/handler"
)

func main(){

	serverBootstrap:=bootstrap.NewServerBootstrap()
	serverBootstrap.Channel(channel.NewTCPServerSocketChannel()).Handler(handler.NewChannelInitializerHandler(func(channel channel_.IChannel){
		fmt.Println("hello")
	})).ChildHandler(handler.NewChannelInitializerHandler(func(channel channel_.IChannel){
		fmt.Println("hello child")
	})).Bind(1024).Sync()



}