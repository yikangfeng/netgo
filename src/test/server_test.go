package test


import (

	"fmt"

	"bootstrap"

	"kernel/channel"
	"kernel/handler"
	"testing"
)

func ServerTest(t *testing.T) {

   serverBootstrap:=bootstrap.NewServerBootstrap()
serverBootstrap.Channel(channel.NewTCPServerSocketChannel()).Handler(handler.NewChannelInitializerHandler(func(channel channel.IChannel){
  fmt.Println(channel)
})).Bind(2014).Sync()

}

