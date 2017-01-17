package main

import (
	"bootstrap"
	channel_ "kernel/intf/external/channel"
	"fmt"
	"kernel/handler"
	"kernel/channel"
)

func main(){

	serverBootstrap:=bootstrap.NewServerBootstrap()
	serverBootstrap.Channel(channel.NewTCPServerSocketChannel()).Handler(handler.NewChannelInitializerHandler(func(channel channel_.IChannel){
		fmt.Println(channel)
	})).Bind(1024).Sync()

}