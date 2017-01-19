package bootstrap

import (
	"kernel/intf/external/handler"
	"kernel/channel"
	"kernel/intf/external/common"

	"log"
)

type ServerBootstrapAcceptor struct {
	//impl IChannelInboundHandler
	handler.IChannelInboundHandler
	ChildOption  map[string]interface{}
	ChildHandler common.IChannelHandler
}

func (this *ServerBootstrapAcceptor) ChannelActive_(ctx handler.IChannelHandlerContext) {
	//do nothing.
}

func (this *ServerBootstrapAcceptor) ChannelInactive_(ctx handler.IChannelHandlerContext) {
	//do nothing.
}

func (this *ServerBootstrapAcceptor) ChannelRead_(ctx handler.IChannelHandlerContext, msg interface{}) {
	if msg == nil {
		return
	}
	socketChannel := msg.(*channel.SocketChannel)
	socketChannel.Config(this.ChildOption)
	if (this.ChildHandler != nil) {
		socketChannel.Pipeline().AddLast(this.ChildHandler)
	}
        go func() {//worker go
		log.Println("start channel read...")
		socketChannel.Pipeline().FireChannelActive()
	}()

}