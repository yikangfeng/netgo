package bootstrap

import (
	"sync"
	"kernel/intf/external/common"
)

/*
@author YiKangfeng.
 */
type ServerBootstrap struct {
	//extends abstractbootstrap
	abstractbootstrap

	childOption  map[string]interface{}

	childHandler common.IChannelHandler
}

var _swait sync.WaitGroup

func NewServerBootstrap() *ServerBootstrap {
	serverBootstrap := &ServerBootstrap{}
	serverBootstrap.option = make(map[string]interface{})
	serverBootstrap.childOption = make(map[string]interface{})
	return serverBootstrap
}

func (this *ServerBootstrap)Option(key  string, v interface{}) *ServerBootstrap {
	_, ok := this.option[key]
	if (!ok) {
		this.option[key] = v
	}
	return this
}

func (this *ServerBootstrap)Channel(_channel common.IChannel) *ServerBootstrap {
	if (_channel == nil) {
		return nil
	}
	this.channel = _channel;
	return this
}

//for Logging Handler
func (this *ServerBootstrap)Handler(channelHandler common.IChannelHandler) *ServerBootstrap {
	this.handler = channelHandler
	return this
}

func (this *ServerBootstrap)Bind(port int) *ServerBootstrap {
	this.BindByHostAndPort("", port)
	return this
}

func (this *ServerBootstrap)BindByHostAndPort(host string, port int) *ServerBootstrap {
	this.internalInit()
	go func() {
		this.channel.Bind(host, port)
	}()

	return this
}

func (this *ServerBootstrap)internalInit() {
	this.init()
	this.channel.Pipeline().AddLast(&ServerBootstrapAcceptor{ChildOption:this.childOption,ChildHandler:this.childHandler})
}

func (this *ServerBootstrap)ChildOption(key  string, v interface{}) *ServerBootstrap {
	_, ok := this.childOption[key]
	if (!ok) {
		this.childOption[key] = v
	}
	return this
}

func (this *ServerBootstrap)ChildHandler(childChannelHandler common.IChannelHandler) *ServerBootstrap {
	this.childHandler = childChannelHandler
	return this
}

func (this *ServerBootstrap)Sync() *ServerBootstrap {

	_swait.Add(1)

	_swait.Wait()

	return this
}
