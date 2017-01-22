package handler

import (
	"kernel/intf/external/common"
)

type IChannelInboundHandler interface {
	//extends IChannelHandler
	ChannelActive_(ctx common.IChannelHandlerContext)
	ChannelInactive_(ctx common.IChannelHandlerContext)
	ChannelRead_(ctx common.IChannelHandlerContext, msg interface{})
}