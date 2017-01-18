package handler

import (
	"kernel/intf/external/handler"
	"kernel/intf/external/channel"
)

type HeadContext struct {
	//impl IChannelOutboundHandler
	AbstractChannelHandlerContext
}

func (this *HeadContext) Connect(host string, port int) {
	this.Channel.(channel.IClientSocketChannel).ConnectAndInit(host, port)
}

func (this *HeadContext) Bind(host string, port int) {
	if _, ok := this.Channel.(channel.IServerSocketChannel); ok {
		this.Channel.(channel.IServerSocketChannel).DoBindAndAccept(host, port)
	} else if _, ok := this.Channel.(channel.IClientSocketChannel); ok {
		//do client bind operation.
	}

}

func (this *HeadContext) Handler() (handler.IChannelHandler) {
	return this
}
