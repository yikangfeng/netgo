package handler

import "kernel/intf/external/common"

type IChannelOutboundHandler interface {
	//extends IChannelHandler
	Connect_(ctx common.IChannelHandlerContext, host string, port int)
	Disconnect_(ctx common.IChannelHandlerContext)
	Bind_(ctx common.IChannelHandlerContext, host string, port int)
	Read_(ctx common.IChannelHandlerContext)
	Close_(ctx common.IChannelHandlerContext)
	Write_(ctx common.IChannelHandlerContext, msg interface{})
	Flush_(ctx common.IChannelHandlerContext)
}


