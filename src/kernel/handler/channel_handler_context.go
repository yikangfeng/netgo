package handler

import "net"

type ChannelHandlerContext struct {
      channel net.Conn
}

func NewChannelHandlerContext(_channel net.Conn) *ChannelHandlerContext {
	return &ChannelHandlerContext{channel:_channel}
}
