package handler

import (
	"kernel/intf/external/common"
)

type ChannelInitializerHandler struct {
	//impl IChannelInboundHandler
	ChannelHandlerAdapter
	ChannelInitFunc func(channel common.IChannel)
}

func NewChannelInitializerHandler(channelInitFunc func(channel common.IChannel)) common.IChannelHandler {
	return &ChannelInitializerHandler{ChannelInitFunc:channelInitFunc}
}

func (this *ChannelInitializerHandler)ChannelActive_(ctx common.IChannelHandlerContext) {
	ctx.FireChannelActive()
}

func (this *ChannelInitializerHandler)ChannelInactive_(ctx common.IChannelHandlerContext) {
	ctx.FireChannelInactive()
}

func (this *ChannelInitializerHandler)ChannelRead_(ctx common.IChannelHandlerContext, msg interface{}) {
	ctx.FireChannelRead(msg)
}