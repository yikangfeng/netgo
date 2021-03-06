package main

import (
	"fmt"
	channel_ "kernel/intf/external/common"
	"bootstrap"
	"kernel/channel"
	"kernel/handler"
	handler_ "kernel/intf/external/handler"
	"sync"
	"bytes"
	"encoding/binary"
)

type test interface {
	add()
}

type test1 struct {
	//impl test
}

type test2 struct {
	test1
}

type TestClientChannelHandler struct {
	//impl IChannelInboundHandler
	handler_.IChannelInboundHandler
}

var _ctx channel_.IChannelHandlerContext

func (this *TestClientChannelHandler)ChannelActive_(ctx channel_.IChannelHandlerContext) {
	fmt.Println("TestClientChannelHandler channel active called.")
	_ctx = ctx
	_wait.Done()
}

func (this *TestClientChannelHandler)ChannelInactive_(ctx channel_.IChannelHandlerContext) {

}

func (this *TestClientChannelHandler)ChannelRead_(ctx channel_.IChannelHandlerContext, msg interface{}) {

}

func (this *TestClientChannelHandler)ExceptionCaught(ctx channel_.IChannelHandlerContext, err error) {
	fmt.Println("has error")
	fmt.Println(err)
}

var _wait sync.WaitGroup

func main() {
	bootstrap := bootstrap.NewBootstrap()
	bootstrap.Channel(channel.NewTCPSocketChannel()).Handler(handler.NewChannelInitializerHandler(func(channel channel_.IChannel) {
		channel.Pipeline().AddLast(&TestClientChannelHandler{})
	})).Connect("127.0.0.1", 1024)

	_wait.Add(1)
	_wait.Wait()

	var network bytes.Buffer

	packetBody := []byte("hello server")
	packetHeaderContent := []byte("protocolName:stringprotocol")

	binary.Write(&network, binary.BigEndian, uint32(len(packetHeaderContent) + len(packetBody) + 4 + 4))

	binary.Write(&network, binary.BigEndian, uint32(len(packetHeaderContent)))

	binary.Write(&network, binary.BigEndian, packetHeaderContent)

	binary.Write(&network, binary.BigEndian, packetBody)

	fmt.Println(uint32(len(network.Bytes())))

	_ctx.Write(network.Bytes())

	var a string = "b"
	var b string = "a"
	if a != b {
		fmt.Println("true")
	}

	//_wait.Add(1)
	//_wait.Wait()

}

func (this *test1) add() {

}