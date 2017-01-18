package channel

import (
	"kernel/intf/external/channel"
	"kernel/intf/external/handler"
	handler_   "kernel/handler"
	"reflect"
	//"fmt"
	//"fmt"
)
//this is DefaultChannelPipeline
type ChannelPipeline struct {
	//impl IChannelPipeline
	channel.IChannelPipeline
	head     handler.IChannelHandlerContext
	tail     handler.IChannelHandlerContext
	channel  channel.IChannel
	handlers map[string]handler.IChannelHandlerContext
}

func NewChannelPipeline(_channel channel.IChannel) channel.IChannelPipeline {
	channelPipeline := &ChannelPipeline{}
	channelPipeline.channel = _channel
	channelPipeline.handlers = make(map[string]handler.IChannelHandlerContext, 4)

	headerCtx := &handler_.HeadContext{}
	headerCtx.SetPrev(nil)
	headerCtx.SetNext(nil)
	headerCtx.Pipeline = channelPipeline
	headerCtx.Channel = _channel
	headerCtx.Name = "headerctx"
	headerCtx.Inbound = false
	headerCtx.Outbound = true

	tailCtx := &handler_.TailContext{}
	tailCtx.SetPrev(nil)
	tailCtx.SetNext(nil)
	tailCtx.Pipeline = channelPipeline
	tailCtx.Channel = _channel
	tailCtx.Name = "tailctx"
	tailCtx.Inbound = true
	tailCtx.Outbound = false

	channelPipeline.head = headerCtx
	channelPipeline.tail = tailCtx

	channelPipeline.head.SetNext(channelPipeline.tail)
	channelPipeline.tail.SetPrev(channelPipeline.head)
	return channelPipeline
}

func (this *ChannelPipeline)AddFirst(handlers ...handler.IChannelHandler) (channel.IChannelPipeline) {
	for i := 0; i < len(handlers); i++ {
		name := reflect.TypeOf(handlers[i]).Name()
		handler := handlers[i]
		newCtx := handler_.NewChannelHandlerContext(this, name, handler)
		this.addFirst(name, newCtx)
	}
	return this
}

func (this *ChannelPipeline)addFirst(name string, newCtx handler.IChannelHandlerContext) {
	if _, ok := this.handlers[name]; !ok {
		nextCtx := this.head.Next()
		newCtx.SetPrev(this.head)
		newCtx.SetNext(nextCtx)
		this.head.SetNext(newCtx)
		nextCtx.SetPrev(newCtx)

		this.handlers[name] = newCtx
	}
}

func (this *ChannelPipeline)AddLast(_handlers ...handler.IChannelHandler) (channel.IChannelPipeline) {
	for i := 0; i < len(_handlers); i++ {
		handler := _handlers[i]
		name := reflect.TypeOf(handler).Name()
		newCtx := handler_.NewChannelHandlerContext(this, name, handler)
		this.addLast(name, newCtx)
	}
	return this
}

func (this *ChannelPipeline)addLast(name string, newCtx handler.IChannelHandlerContext) {
	if _, ok := this.handlers[name]; !ok {
		_prev := this.tail.Prev();
		newCtx.SetPrev(_prev);
		newCtx.SetNext(this.tail);
		_prev.SetNext(newCtx);
		this.tail.SetPrev(newCtx);
		this.handlers[name] = newCtx
	}
}

func (this *ChannelPipeline)AddBefore(baseName string, _handlers ...handler.IChannelHandler) (channel.IChannelPipeline) {
	for i := 0; i < len(_handlers); i++ {
		handler := _handlers[i]
		name := reflect.TypeOf(handler).Name()
		if _, ok := this.handlers[baseName]; !ok {
			return this
		}
		if _, ok := this.handlers[name]; !ok {
			ctx := this.handlers[baseName]
			newCtx := handler_.NewChannelHandlerContext(this, name, handler)
			this.addBefore(name, ctx, newCtx)
		}
	}
	return this
}

func (this *ChannelPipeline)addBefore(name string, ctx handler.IChannelHandlerContext, newCtx handler.IChannelHandlerContext) {

	if _, ok := this.handlers[name]; !ok {
		newCtx.SetPrev(ctx.Prev())
		newCtx.SetNext(ctx)
		ctx.Prev().SetNext(newCtx)
		ctx.SetPrev(newCtx)
		this.handlers[name] = newCtx
	}
}

func (this *ChannelPipeline)AddAfter(baseName string, _handlers ...handler.IChannelHandler) (channel.IChannelPipeline) {
	for i := 0; i < len(_handlers); i++ {
		handler := _handlers[i]
		name := reflect.TypeOf(handler).Name()
		if _, ok := this.handlers[baseName]; !ok {
			return this
		}
		if _, ok := this.handlers[name]; !ok {
			ctx := this.handlers[baseName]
			newCtx := handler_.NewChannelHandlerContext(this, name, handler)
			this.addAfter(name, ctx, newCtx)
		}
	}
	return this
}

func (this *ChannelPipeline)addAfter(name string, ctx handler.IChannelHandlerContext, newCtx handler.IChannelHandlerContext) {

	if _, ok := this.handlers[name]; !ok {
		newCtx.SetPrev(ctx)
		newCtx.SetNext(ctx.Next())
		ctx.Next().SetPrev(newCtx)
		ctx.SetNext(newCtx);

		this.handlers[name] = newCtx
	}

}

func (this *ChannelPipeline)RemoveFirst() (handler.IChannelHandler) {

	if this.head.Next() == this.tail {
		return nil
	}

	return this.remove(this.head.Next()).Handler()

}

func (this *ChannelPipeline)RemoveLast() (handler.IChannelHandler) {

	if this.head.Next() == this.tail {
		return nil
	}

	return this.remove(this.tail.Prev()).Handler()

}

func (this *ChannelPipeline)remove(ctx handler.IChannelHandlerContext) (handler.IChannelHandlerContext) {
	if ctx == nil {
		return nil
	}
	prev := ctx.Prev();
	next := ctx.Next();
	prev.SetNext(next);
	next.SetPrev(prev);
	deletedCtx := this.handlers[ctx.GetName()]
	delete(this.handlers, ctx.GetName())
	return deletedCtx
}

func (this *ChannelPipeline)Remove(handler handler.IChannelHandler) {
	if handler == nil {
		return
	}
	name := reflect.TypeOf(handler).Name()
	if _, ok := this.handlers[name]; ok {
		this.remove(this.handlers[name])
	}
}

func (this *ChannelPipeline)Replace(oldHandler handler.IChannelHandler, newName string, newHandler handler.IChannelHandler) (channel.IChannelPipeline) {
	if oldHandler == nil {
		return this
	}
	ctx := this.context(oldHandler)
	if ctx == nil {
		return this
	}

	var sameName bool = (ctx.GetName() == newName)
	if (!sameName) {
		if _, ok := this.handlers[newName]; ok {
			//newName exists.
			return this
		}
	}
	newCtx := handler_.NewChannelHandlerContext(this, newName, newHandler)
	this.replace(ctx, newName, newCtx)
	return this
}

func (this *ChannelPipeline)replace(oldCtx handler.IChannelHandlerContext, newName string, newCtx handler.IChannelHandlerContext) {
	prev := oldCtx.Prev()
	next := oldCtx.Next()
	newCtx.SetPrev(prev)
	newCtx.SetNext(next)

	prev.SetNext(newCtx)
	next.SetPrev(newCtx)

	if (oldCtx.GetName() != newName) {
		delete(this.handlers, oldCtx.GetName())
	}
	this.handlers[newName] = newCtx

	oldCtx.SetPrev(newCtx)
	oldCtx.SetNext(newCtx)
}

func (this *ChannelPipeline)context(handler handler.IChannelHandler) (handler.IChannelHandlerContext) {
	if handler == nil {
		return nil
	}

	ctx := this.head.Next();
	for {

		if (ctx == nil) {
			return nil;
		}

		if (ctx.Handler() == handler) {
			return ctx;
		}

		ctx = ctx.Next();
	}
}

func (this *ChannelPipeline)Connect(host string, port int) {
	this.tail.(*handler_.TailContext).AbstractChannelHandlerContext.Connect(host, port)
}

func (this *ChannelPipeline)Bind(host string, port int) {
	this.tail.(*handler_.TailContext).AbstractChannelHandlerContext.Bind(host, port)
}

func (this *ChannelPipeline)GetChannel() channel.IChannel {
	return this.channel
}