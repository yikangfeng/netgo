package channel

import (
	"kernel/intf/external/handler"
)

type IChannelPipeline interface {
	AddFirst(handlers ...handler.IChannelHandler) (IChannelPipeline)

	AddBefore(baseName string, handlers ...handler.IChannelHandler) (IChannelPipeline)
	AddAfter(baseName string, handlers ...handler.IChannelHandler) (IChannelPipeline)

	AddLast(_handlers ...handler.IChannelHandler) (IChannelPipeline)

	RemoveFirst() (handler.IChannelHandler)
	RemoveLast() (handler.IChannelHandler)
	Remove(handler handler.IChannelHandler)

	replace(oldHandler handler.IChannelHandler, newName string, newHandler handler.IChannelHandler) (IChannelPipeline)

	Connect(host string, port int)
	Bind(host string, port int)
	GetChannel() IChannel
}
