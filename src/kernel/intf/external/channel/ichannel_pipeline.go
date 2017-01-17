package channel

import (

	"kernel/intf/external/handler"
)


type IChannelPipeline interface {
Add()

AddFirst(handlers ...handler.IChannelHandler)

AddLast(_handlers ...handler.IChannelHandler)

Connect(host string,port int)
	Bind(host string,port int)
GetChannel() IChannel
}
