package handler

import (
	"net"
	"kernel/channel"
	"context"
)

type AbstractChannelHandlerContext struct {//impl IChannelHandlerContext
	Next *AbstractChannelHandlerContext
	Prev *AbstractChannelHandlerContext
	Pipeline *channel.ChannelPipeline
	Channel  channel.IChannel
	Name string
	Inbound bool
	Outbound bool
}


type ChannelHandlerContext struct {//extends AbstractChannelHandlerContext
	AbstractChannelHandlerContext
	handler IChannelHandler

}

func NewChannelHandlerContext(_pipeline *channel.ChannelPipeline,_name string,_handler *IChannelHandler ) *ChannelHandlerContext {
	return NewChannelHandlerContext0(_pipeline,_name,isInbound(_handler),isOutbound(_handler),_handler)
}

func NewChannelHandlerContext0(_pipeline *channel.ChannelPipeline,_name string,inBound bool,outBound bool,_handler *IChannelHandler ) *ChannelHandlerContext {
	context:=&ChannelHandlerContext{}
	context.Inbound=inBound
	context.Outbound=outBound
	context.Name=_name
	context.Pipeline=_pipeline
	context.Channel=_pipeline.GetChannel()
	context.handler=_handler
	return context
}

func isInbound(_handler IChannelHandler) bool {
	_,ok:= _handler.(IChannelInboundHandler)
	return ok
}

func isOutbound(_handler IChannelHandler) bool {
	_,ok:= _handler.(IChannelOutboundHandler)
	return ok
}

func (this *ChannelHandlerContext)Handler() IChannelHandler {
	return this.handler
}

func (this *AbstractChannelHandlerContext)Connect(host string,port int)  {
	//this.findContextOutbound()
        //this.Handler().(IChannelOutboundHandler).Connect(host,port)

}

func (this *ChannelHandlerContext)findContextInbound() *ChannelHandlerContext {
	ctx := this;
	for {
		ctx = ctx.Next;
		if ctx.Inbound {
			break
		}
	}
	return ctx
}

func (this *ChannelHandlerContext)findContextOutbound() *ChannelHandlerContext {
	ctx := this;
	for {
		ctx = ctx.Prev;
		if ctx.Outbound {
			break
		}
	}
	return ctx
}