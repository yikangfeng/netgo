package handler

import (
	"kernel/intf/external/channel"
	"kernel/intf/external/common"
)

type HeadContext struct {
	//impl IChannelOutboundHandler
	AbstractChannelHandlerContext
}

func (this *HeadContext) Connect_(ctx common.IChannelHandlerContext, host string, port int) {
	this.GetChannel().(channel.IClientSocketChannel).ConnectAndInit(host, port)
}

func (this *HeadContext) Bind_(ctx common.IChannelHandlerContext, host string, port int) {
	if _, ok := this.GetChannel().(channel.IServerSocketChannel); ok {
		this.GetChannel().(channel.IServerSocketChannel).DoBindAndAccept(host, port)
		go func() {
			this.GetChannel().Pipeline().FireChannelActive()
		}()
	} else if _, ok := this.GetChannel().(channel.IClientSocketChannel); ok {
		//do client bind operation.
	}

}

func (this *HeadContext) Disconnect_(ctx common.IChannelHandlerContext) {

}

func (this *HeadContext) Read_(ctx common.IChannelHandlerContext) {
}
func (this *HeadContext) Close_(ctx common.IChannelHandlerContext) {
	defer this.GetChannel().Close()
}
func (this *HeadContext) Write_(ctx common.IChannelHandlerContext, msg interface{}) {
	this.GetChannel().(channel.IClientSocketChannel).GetConn().Write(msg.([]byte))
}

func (this *HeadContext) Flush_(ctx common.IChannelHandlerContext) {

}

func (this *HeadContext) ExceptionCaught(ctx common.IChannelHandlerContext, err error) {
	ctx.FireExceptionCaught(err)
}

func (this *HeadContext) Handler() (common.IChannelHandler) {
	return this
}
