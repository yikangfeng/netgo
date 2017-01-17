package handler

import (
	"kernel/intf/external/handler"
	"kernel/intf/external/channel"
)
type ChannelInitializerHandler struct {//impl IChannelInboundHandler
     handler.IChannelInboundHandler
     ChannelInitFunc func(channel channel.IChannel)
}

func NewChannelInitializerHandler(channelInitFunc func(channel channel.IChannel))  handler.IChannelHandler{
	return &ChannelInitializerHandler{ChannelInitFunc:channelInitFunc}
}