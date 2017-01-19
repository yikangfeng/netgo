package handler

import "kernel/intf/external/common"

type IChannelHandlerContext interface {
	//handler context
	Connect(host string, port int)
	Bind(host string, port int)
	Handler() common.IChannelHandler
	GetChannel() common.IChannel
	GetName() string
	Next() IChannelHandlerContext
	SetNext(ctx IChannelHandlerContext)
	Prev() IChannelHandlerContext
	SetPrev(ctx IChannelHandlerContext)
	GetInbound() bool
	GetOutbound() bool
	FireChannelActive() (IChannelHandlerContext)
	FireChannelInactive() (IChannelHandlerContext)
	FireChannelRead(msg interface{}) (IChannelHandlerContext)
	Read() (IChannelHandlerContext)
	Write(msg interface{})
	Flush()
	WriteAndFlush(msg interface{})
	Close()
}
