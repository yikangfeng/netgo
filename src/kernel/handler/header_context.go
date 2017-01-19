package handler

import (
	"kernel/intf/external/channel"
	"kernel/intf/external/common"
	"kernel/intf/external/handler"
)

type HeadContext struct {
	//impl IChannelOutboundHandler
	AbstractChannelHandlerContext
}

func (this *HeadContext) Connect_(ctx handler.IChannelHandlerContext, host string, port int) {

	this.GetChannel().(channel.IClientSocketChannel).ConnectAndInit(host, port)
}

func (this *HeadContext) Bind_(ctx handler.IChannelHandlerContext, host string, port int) {
	if _, ok := this.GetChannel().(channel.IServerSocketChannel); ok {
		this.GetChannel().(channel.IServerSocketChannel).DoBindAndAccept(host, port)
		go func() {
			this.GetChannel().Pipeline().FireChannelActive()
		}()
	} else if _, ok := this.GetChannel().(channel.IClientSocketChannel); ok {
		//do client bind operation.
	}

}

func (this *HeadContext) Disconnect_(ctx handler.IChannelHandlerContext) {

}

func (this *HeadContext) Read_(ctx handler.IChannelHandlerContext) {
}
func (this *HeadContext) Close_(ctx handler.IChannelHandlerContext) {

}
func (this *HeadContext) Write_(ctx handler.IChannelHandlerContext, msg interface{}) {

}

func (this *HeadContext) Flush_(ctx handler.IChannelHandlerContext) {

}

func (this *HeadContext) Handler() (common.IChannelHandler) {
	return this
}
