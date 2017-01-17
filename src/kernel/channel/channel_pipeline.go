package channel

import (
     "kernel/intf/external/channel"
     "kernel/intf/external/handler"
      handler_   "kernel/handler"
     "reflect"
)
//this is DefaultChannelPipeline
type ChannelPipeline struct {//impl IChannelPipeline
     channel.IChannelPipeline
     head handler.IChannelHandlerContext
     tail handler.IChannelHandlerContext
     channel channel.IChannel
     handlers map[string]handler.IChannelHandlerContext
}


type TailContext struct {//impl IChannelInboundHandler
     handler_.AbstractChannelHandlerContext
}

func (this *TailContext) Bind(host string,port int) {
     this.Channel.(IServerSocketChannel).doBindAndAccept(host,port)
}

func (this *TailContext) Connect(host string,port int) {

}

func (this *TailContext) Handler() (handler.IChannelHandler){
     return this
}

type HeadContext struct {//impl IChannelOutboundHandler
     handler_.AbstractChannelHandlerContext

}

func (this *HeadContext) Connect(host string,port int) {
     this.Channel.(IClientSocketChannel).connectAndInit(host,port)
}

func (this *HeadContext) Bind(host string,port int) {

}

func (this *HeadContext) Handler() (handler.IChannelHandler){
      return this
}


func NewChannelPipeline(_channel channel.IChannel)  *ChannelPipeline{
     channelPipeline:=&ChannelPipeline{}
     channelPipeline.channel=_channel
     channelPipeline.handlers=make(map[string]handler.IChannelHandlerContext,4)

     headerCtx:=&HeadContext{}
     headerCtx.SetPrev(nil)
     headerCtx.SetNext(nil)
     headerCtx.Pipeline=channelPipeline
     headerCtx.Channel=_channel
     headerCtx.Name="headerctx"
     headerCtx.Inbound=false
     headerCtx.Outbound=true

     tailCtx:=&TailContext{}
     tailCtx.SetPrev(nil)
     tailCtx.SetNext(nil)
     tailCtx.Pipeline=channelPipeline
     tailCtx.Channel=_channel
     tailCtx.Name="tailctx"
     tailCtx.Inbound=true
     tailCtx.Outbound=false

     channelPipeline.head=headerCtx
     channelPipeline.tail=tailCtx

     channelPipeline.head.SetNext(channelPipeline.tail)
     channelPipeline.tail.SetPrev(channelPipeline.head)
     return channelPipeline
}

func (this *ChannelPipeline)Add() {

}

func (this *ChannelPipeline)AddFirst(handlers ...handler.IChannelHandler) {

}

func (this *ChannelPipeline)AddLast(_handlers ...handler.IChannelHandler) {
     for i:=0;i<len(_handlers);i++{
          handlerInstance:=_handlers[i]
          if _,ok:= this.handlers[reflect.TypeOf(handlerInstance).Name()];!ok {
               newCtx:=handler_.NewChannelHandlerContext(this, reflect.TypeOf(handlerInstance).Name(),handlerInstance)

               _prev := this.tail.Prev();
               newCtx.SetPrev(_prev);
               newCtx.SetNext(this.tail);
               _prev.SetNext(newCtx);
               this.tail.SetPrev(newCtx);
               this.handlers[reflect.TypeOf(handlerInstance).Name()]=newCtx
          }
     }
}

func (this *ChannelPipeline)Connect(host string,port int) {
     this.tail.Connect(host,port)
}


func (this *ChannelPipeline)Bind(host string,port int) {
     this.tail.Bind(host,port)
}

func (this *ChannelPipeline)GetChannel() channel.IChannel{
    return this.channel
}