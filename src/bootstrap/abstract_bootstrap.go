package bootstrap

import (
	"kernel/intf/external/common"
	handler_ "kernel/handler"
)

type abstractbootstrap struct {
	channel common.IChannel
	handler common.IChannelHandler//Initialization handler.
	option  map[string]interface{}
}

func (this *abstractbootstrap)init() {
	this.channel.(common.IChannel).Config(this.option)
	this.handler.(*handler_.ChannelInitializerHandler).ChannelInitFunc(this.channel)
}
