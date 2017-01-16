package client

import (
"sync"
	"kernel/handler"
	"kernel/channel"
)

/*
@author YiKangfeng.
 */
type ClientBootstrap struct {

	option map[string]interface{}

	channel channel.IChannel

	handler handler.IChannelHandler

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



func (this *ClientBootstrap)Handler( channelHandler handler.IChannelHandler) *ClientBootstrap {
	this.handler=channelHandler
	return this
}

func (this *ClientBootstrap)Connect(host string,port int) *ClientBootstrap {
        this.init()
	clientSocketChannel:=this.channel.(channel.AbstractSocketChannel)
	go func() {
		clientSocketChannel.Connect(host, port)
	}()
	return this
}

func (this *ClientBootstrap)init()  {
	this.channel.(channel.AbstractSocketChannel).Config(&this.option)
     this.channel.Pipeline().AddLast(this.handler)

}

func (this *ClientBootstrap)Sync() *ClientBootstrap {
	return this
}