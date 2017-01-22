package handler

import (
	"kernel/intf/external/common"
)

type ChannelHandlerAdapter struct {
	//impl IChannelHandler
}

func (this *ChannelHandlerAdapter)ExceptionCaught(ctx common.IChannelHandlerContext, err error) {
	ctx.FireExceptionCaught(err)
}