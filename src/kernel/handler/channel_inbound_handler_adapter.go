package handler

import "kernel/intf/external/common"

type ChannelInboundHandlerAdapter struct {
	//impl IChannelInboundHandler
}

func (this *ChannelInboundHandlerAdapter)ChannelActive_(ctx common.IChannelHandlerContext) {

}

func (this *ChannelInboundHandlerAdapter)ChannelInactive_(ctx common.IChannelHandlerContext) {

}

func (this *ChannelInboundHandlerAdapter)ExceptionCaught(ctx common.IChannelHandlerContext, err error) {

}




