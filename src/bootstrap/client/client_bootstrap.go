package client

import (
"sync"
	"kernel"
	"kernel/handler"
	"net"
	"fmt"
	"os"
	"kernel/channel"
	"kernel/channel/socket"
)

/*
@author YiKangfeng.
 */
type ClientBootstrap struct {

	option map[string]interface{}

	channel channel.IChannel


}

var _wait sync.WaitGroup

func  New() *ClientBootstrap{
	return &ClientBootstrap{option:make(map[string]interface{})}
}

func (this *ClientBootstrap)Channel( _channel *channel.IChannel) (_this *ClientBootstrap) {
	if(_channel==nil){
		return nil
	}
	this.channel=_channel;
	return this
}

func (this *ClientBootstrap)Option( key  string, v interface{}) *ClientBootstrap {
	_,ok := this.option[key]
	if(!ok) {
		this.option[key] = v
	}
	return this
}



func (this *ClientBootstrap)Handler( channelHandler handler.ChannelHandler) *ClientBootstrap {
	return this
}

func (this *ClientBootstrap)Connect(host string,port int) *ClientBootstrap {
	this.channel.(socket.IClientSocketChannel).Connect(host,port)
	return this
}
func (this *ClientBootstrap)Sync() *ClientBootstrap {
	return this
}