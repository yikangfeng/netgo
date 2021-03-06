package handler

import (
	"kernel/intf/external/handler"
	"kernel/intf/external/common"
)

type AbstractChannelHandlerContext struct {
	//impl IChannelHandlerContext
	common.IChannelHandlerContext
	next     common.IChannelHandlerContext
	prev     common.IChannelHandlerContext
	Pipeline common.IChannelPipeline
	Channel  common.IChannel
	Name     string
	Inbound  bool
	Outbound bool
}

type ChannelHandlerContext struct {
	//extends AbstractChannelHandlerContext
	AbstractChannelHandlerContext
	handler common.IChannelHandler
}

func NewChannelHandlerContext(_pipeline common.IChannelPipeline, _name string, _handler common.IChannelHandler) common.IChannelHandlerContext {
	return NewChannelHandlerContext0(_pipeline, _name, isInbound(_handler), isOutbound(_handler), _handler)
}

func NewChannelHandlerContext0(_pipeline common.IChannelPipeline, _name string, inBound bool, outBound bool, _handler common.IChannelHandler) common.IChannelHandlerContext {
	context := &ChannelHandlerContext{}
	context.Inbound = inBound
	context.Outbound = outBound
	context.Name = _name
	context.Pipeline = _pipeline
	context.Channel = _pipeline.GetChannel()
	context.handler = _handler
	return context
}

func isInbound(_handler common.IChannelHandler) bool {
	_, ok := _handler.(handler.IChannelInboundHandler)
	return ok
}

func isOutbound(_handler common.IChannelHandler) bool {
	_, ok := _handler.(handler.IChannelOutboundHandler)
	return ok
}

func (this *ChannelHandlerContext)Handler() common.IChannelHandler {
	return this.handler
}

func (this *AbstractChannelHandlerContext)GetChannel() common.IChannel {
	return this.Channel
}

func (this *AbstractChannelHandlerContext)GetPipeline() common.IChannelPipeline {
	return this.Pipeline
}

func (this *AbstractChannelHandlerContext)Next() (common.IChannelHandlerContext) {
	return this.next
}

func (this *AbstractChannelHandlerContext)Prev() (common.IChannelHandlerContext) {
	return this.prev
}

func (this *AbstractChannelHandlerContext)SetNext(ctx common.IChannelHandlerContext) {
	this.next = ctx;
}

func (this *AbstractChannelHandlerContext)SetPrev(ctx common.IChannelHandlerContext) {
	this.prev = ctx
}

func (this *AbstractChannelHandlerContext)GetName() string {
	return this.Name
}

func (this *AbstractChannelHandlerContext)GetInbound() (bool) {
	return this.Inbound
}

func (this *AbstractChannelHandlerContext)GetOutbound() (bool) {
	return this.Outbound
}

func (this *AbstractChannelHandlerContext)findContextInbound() common.IChannelHandlerContext {
	var ctx common.IChannelHandlerContext = this
	for {
		ctx = ctx.Next();
		if ctx.GetInbound() {
			break
		}
	}
	return ctx
}

func (this *AbstractChannelHandlerContext)findContextOutbound() common.IChannelHandlerContext {
	var ctx common.IChannelHandlerContext = this
	for {
		ctx = ctx.Prev();
		if ctx.GetOutbound() {
			break
		}
	}
	return ctx
}

func (this *AbstractChannelHandlerContext)Connect(host string, port int) {
	ctx := this.findContextOutbound()
	ctx.Handler().(handler.IChannelOutboundHandler).Connect_(ctx, host, port)
}

func (this *AbstractChannelHandlerContext)Bind(host string, port int) {
	ctx := this.findContextOutbound()
	ctx.Handler().(handler.IChannelOutboundHandler).Bind_(ctx, host, port)
}

func (this *AbstractChannelHandlerContext)FireChannelActive() (common.IChannelHandlerContext) {
	ctx := this.findContextInbound()
	ctx.Handler().(handler.IChannelInboundHandler).ChannelActive_(ctx)
	return this
}

func (this *AbstractChannelHandlerContext) FireChannelRead(msg interface{}) (common.IChannelHandlerContext) {
	ctx := this.findContextInbound()
	ctx.Handler().(handler.IChannelInboundHandler).ChannelRead_(ctx, msg)
	return this
}

func (this *AbstractChannelHandlerContext) FireChannelInactive() (common.IChannelHandlerContext) {
	ctx := this.findContextInbound()
	ctx.Handler().(handler.IChannelInboundHandler).ChannelInactive_(ctx)
	return this
}

func (this *AbstractChannelHandlerContext) WriteAndFlush(msg interface{}) {
	ctx := this.findContextOutbound()
	ctx.Handler().(handler.IChannelOutboundHandler).Write_(ctx, msg)
}

func (this *AbstractChannelHandlerContext) Write(msg interface{}) {
	ctx := this.findContextOutbound()
	ctx.Handler().(handler.IChannelOutboundHandler).Write_(ctx, msg)
}

func (this *AbstractChannelHandlerContext) Flush() {
	ctx := this.findContextOutbound()
	ctx.Handler().(handler.IChannelOutboundHandler).Flush_(ctx)
}

func (this *AbstractChannelHandlerContext) Close() {
	ctx := this.findContextOutbound()
	ctx.Handler().(handler.IChannelOutboundHandler).Close_(ctx)
}

func (this *AbstractChannelHandlerContext) Read() (common.IChannelHandlerContext) {
	ctx := this.findContextOutbound()
	ctx.Handler().(handler.IChannelOutboundHandler).Read_(ctx)
	return this
}

func (this *AbstractChannelHandlerContext) FireExceptionCaught(err error) (common.IChannelHandlerContext) {
	ctx := this.Next()
	ctx.Handler().ExceptionCaught(ctx, err)
	return this
}
