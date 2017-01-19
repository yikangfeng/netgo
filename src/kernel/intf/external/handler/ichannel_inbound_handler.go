package handler

import "kernel/intf/external/common"

type IChannelInboundHandler interface {
	//impl IChannelHandler
	common.IChannelHandler
	ChannelActive_(ctx IChannelHandlerContext)
	ChannelInactive_(ctx IChannelHandlerContext)
	ChannelRead_(ctx IChannelHandlerContext, msg interface{})
}