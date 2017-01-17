package bootstrap

import (
	"kernel/intf/external/channel"
	"kernel/intf/external/handler"
//	channel_ "kernel/channel"
//	handler_ "kernel/handler"
)

type abstractbootstrap struct {
	channel channel.IChannel
	handler handler.IChannelHandler
	option map[string]interface{}
}

func (this *abstractbootstrap)init()  {
//	this.channel.(channel_.AbstractSocketChannel).Config(&this.option)
//	this.handler.(handler_.ChannelInitializerHandler).ChannelInitFunc(this.channel)
}
