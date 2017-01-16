package channel

import (
     "kernel/handler"
     "reflect"
)

type ChannelPipeline struct {
     head *handler.AbstractChannelHandlerContext
     tail *handler.AbstractChannelHandlerContext
     channel IChannel
     handlers map[string]handler.AbstractChannelHandlerContext
}


type TailContext struct {//impl IChannelInboundHandler
     handler.AbstractChannelHandlerContext
}

type HeadContext struct {//impl IChannelOutboundHandler
     handler.AbstractChannelHandlerContext

}

func (this *HeadContext) Connect(host string,port int) {
     this.Channel.(IClientSocketChannel).connectAndInit(host,port)
}

func NewChannelPipeline(_channel IChannel)  *ChannelPipeline{
     channelPipeline:=&ChannelPipeline{}
     channelPipeline.channel=_channel
     channelPipeline.handlers=make(map[string]handler.ChannelHandlerContext,4)
     headerCtx:=&HeadContext{}
     tailCtx:=&TailContext{}
     channelPipeline.head=headerCtx
     channelPipeline.tail=tailCtx
     channelPipeline.head.Next=channelPipeline.tail
     channelPipeline.tail.Prev=channelPipeline.head
     return channelPipeline
}

func (this *ChannelPipeline)Add() {

}

func (this *ChannelPipeline)AddFirst(handlers ...handler.IChannelHandler) {

}

func (this *ChannelPipeline)AddLast(_handlers ...*handler.IChannelHandler) {
     for i:=0;i<len(_handlers);i++{
          handlerInstance:=_handlers[i]
          if _,ok:= this.handlers[reflect.TypeOf(handlerInstance).Name()];!ok {
               newCtx:=handler.NewChannelHandlerContext(this, reflect.TypeOf(handlerInstance).Name(),handlerInstance)

               _prev := this.tail.Prev;
               newCtx.Prev = _prev;
               newCtx.Next = this.tail;
               _prev.Next = newCtx;
               this.tail.Prev = newCtx;
               _handlers[reflect.TypeOf(handlerInstance).Name()]=newCtx
          }
     }
}

func (this *ChannelPipeline)Connect(host string,port int) {
     this.tail.Connect(host,port)
}

func (this *ChannelPipeline)GetChannel() IChannel{
    return this.channel
}