package bootstrap

import (
	"kernel/intf/external/channel"
	"kernel/intf/external/handler"
	handler_ "kernel/handler"
)

type abstractbootstrap struct {
	channel channel.IChannel
	handler handler.IChannelHandler//Initialization handler.
	option  map[string]interface{}
}

func (this *abstractbootstrap)init() {
	this.channel.(channel.IChannel).Config(this.option)
	this.handler.(*handler_.ChannelInitializerHandler).ChannelInitFunc(this.channel)
}
