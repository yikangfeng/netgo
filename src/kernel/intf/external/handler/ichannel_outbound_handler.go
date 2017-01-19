package handler

import "kernel/intf/external/common"

type IChannelOutboundHandler interface {
	common.IChannelHandler
	Connect_(ctx IChannelHandlerContext, host string, port int)
	Disconnect_(ctx IChannelHandlerContext)
	Bind_(ctx IChannelHandlerContext, host string, port int)
	Read_(ctx IChannelHandlerContext)
	Close_(ctx IChannelHandlerContext)
	Write_(ctx IChannelHandlerContext, msg interface{})
	Flush_(ctx IChannelHandlerContext)
}


