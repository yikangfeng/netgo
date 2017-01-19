package handler

import (
	"kernel/intf/external/handler"
	"kernel/intf/external/common"
)
type ChannelInitializerHandler struct {//impl IChannelInboundHandler
     handler.IChannelInboundHandler
     ChannelInitFunc func(channel common.IChannel)
}

func NewChannelInitializerHandler(channelInitFunc func(channel common.IChannel))  common.IChannelHandler{
	return &ChannelInitializerHandler{ChannelInitFunc:channelInitFunc}
}