package common

type IChannelPipeline interface {
	AddFirst(handlers ...IChannelHandler) (IChannelPipeline)

	AddBefore(baseName string, handlers ...IChannelHandler) (IChannelPipeline)
	AddAfter(baseName string, handlers ...IChannelHandler) (IChannelPipeline)

	AddLast(_handlers ...IChannelHandler) (IChannelPipeline)

	RemoveFirst() (IChannelHandler)
	RemoveLast() (IChannelHandler)
	Remove(handler IChannelHandler)

	Replace(oldHandler IChannelHandler, newName string, newHandler IChannelHandler) (IChannelPipeline)

	Connect(host string, port int)
	Bind(host string, port int)
	GetChannel() IChannel
	FireChannelActive() (IChannelPipeline)
	FireChannelRead(msg interface{}) (IChannelPipeline)
	FireChannelInactive() (IChannelPipeline)
	Read() (IChannelPipeline)
	Write(msg interface{}) (IChannelPipeline)
	Flush() (IChannelPipeline)
	WriteAndFlush(msg interface{}) (IChannelPipeline)
}
