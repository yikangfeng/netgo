package test


import (

	"fmt"

	"bootstrap"

	"kernel/channel"
	"kernel/handler"
	"testing"
)

func ClientTest(t *testing.T) {

	clientBootstrap:=bootstrap.NewBootstrap()
	clientBootstrap.Channel(channel.NewTCPSocketChannel()).Handler(handler.NewChannelInitializerHandler(func(channel channel.IChannel){
		fmt.Println(channel)
	})).Connect("127.0.0.1",1024);


}

