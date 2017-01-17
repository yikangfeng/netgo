package handler

import (
	"kernel/intf/external/handler"
	"kernel/intf/external/channel"
)

type AbstractChannelHandlerContext struct {//impl IChannelHandlerContext
	handler.IChannelHandlerContext
	next handler.IChannelHandlerContext
	prev handler.IChannelHandlerContext
	Pipeline channel.IChannelPipeline
	Channel  channel.IChannel
	Name string
	Inbound bool
	Outbound bool
}


type ChannelHandlerContext struct {//extends AbstractChannelHandlerContext
	AbstractChannelHandlerContext
	handler handler.IChannelHandler

}

func NewChannelHandlerContext(_pipeline channel.IChannelPipeline,_name string,_handler handler.IChannelHandler ) *ChannelHandlerContext {
	return NewChannelHandlerContext0(_pipeline,_name,isInbound(_handler),isOutbound(_handler),_handler)
}

func NewChannelHandlerContext0(_pipeline channel.IChannelPipeline,_name string,inBound bool,outBound bool,_handler handler.IChannelHandler ) *ChannelHandlerContext {
	context:=&ChannelHandlerContext{}
	context.Inbound=inBound
	context.Outbound=outBound
	context.Name=_name
	context.Pipeline=_pipeline
	context.Channel=_pipeline.GetChannel()
	context.handler=_handler
	return context
}

func isInbound(_handler handler.IChannelHandler) bool {
	_,ok:= _handler.(handler.IChannelInboundHandler)
	return ok
}

func isOutbound(_handler handler.IChannelHandler) bool {
	_,ok:= _handler.(handler.IChannelOutboundHandler)
	return ok
}

func (this *ChannelHandlerContext)Handler() handler.IChannelHandler {
	return this.handler
}


func (this *AbstractChannelHandlerContext)Next() (handler.IChannelHandlerContext)  {
	return this.next
}

func (this *AbstractChannelHandlerContext)Prev() (handler.IChannelHandlerContext)  {
	return this.prev
}

func (this *AbstractChannelHandlerContext)SetNext(ctx handler.IChannelHandlerContext)  {
	this.next=ctx;
}

func (this *AbstractChannelHandlerContext)SetPrev(ctx handler.IChannelHandlerContext) {
	this.prev=ctx
}


func (this *AbstractChannelHandlerContext)Connect(host string,port int)  {
	this.findContextOutbound().Connect(host,port)
}

func (this *AbstractChannelHandlerContext)Bind(host string,port int)  {
	this.findContextOutbound().Bind(host,port)
}

func (this *AbstractChannelHandlerContext)GetInbound() (bool)  {
	return this.Inbound
}

func (this *AbstractChannelHandlerContext)GetOutbound() (bool)  {
	return this.Outbound
}


func (this *AbstractChannelHandlerContext)Handler() (handler.IChannelHandler)  {
	return this
}

func (this *AbstractChannelHandlerContext)findContextInbound() handler.IChannelHandlerContext {
	var ctx handler.IChannelHandlerContext;
	for {
		ctx = ctx.Next();
		if ctx.GetInbound() {
			break
		}
	}
	return ctx
}

func (this *AbstractChannelHandlerContext)findContextOutbound() handler.IChannelHandlerContext {
	var ctx handler.IChannelHandlerContext;
	for {
		ctx = ctx.Prev();
		if ctx.GetOutbound() {
			break
		}
	}
	return ctx
}