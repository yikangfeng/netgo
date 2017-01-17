package bootstrap

import (
	"sync"
	"kernel/intf/external/channel"
	"kernel/intf/external/handler"
)

/*
@author YiKangfeng.
 */
type ServerBootstrap struct {//extends abstractbootstrap
	abstractbootstrap

 childOption map[string]interface{}

 childHandler handler.IChannelHandler

}

var _swait sync.WaitGroup

func  NewServerBootstrap() *ServerBootstrap{
	serverBootstrap:= &ServerBootstrap{}
	serverBootstrap.option=make(map[string]interface{})
	serverBootstrap.childOption=make(map[string]interface{})
	return serverBootstrap
}

func (this *ServerBootstrap)Option( key  string, v interface{}) *ServerBootstrap {
	_,ok := this.option[key]
	if(!ok) {
		this.option[key] = v
	}
	return this
}


func (this *ServerBootstrap)Channel( _channel channel.IChannel) *ServerBootstrap {
	if(_channel==nil){
		return nil
	}
	this.channel=_channel;
	return this
}

//for Logging Handler
func (this *ServerBootstrap)Handler( channelHandler handler.IChannelHandler) *ServerBootstrap {
	this.handler=channelHandler
	return this
}

func (this *ServerBootstrap)Bind(port int) *ServerBootstrap {
	this.BindByHostAndPort("",port)
	return this
}

func (this *ServerBootstrap)BindByHostAndPort(host string,port int) *ServerBootstrap {
	this.init()

	go func() {
		this.channel.Bind(host, port)
	}()

	return this
}


func (this *ServerBootstrap)ChildOption( key  string, v interface{}) *ServerBootstrap {
	_,ok := this.option[key]
	if(!ok) {
		this.option[key] = v
	}
        return this
}

func (this *ServerBootstrap)ChildHandler( childChannelHandler handler.IChannelHandler) *ServerBootstrap {
	this.childHandler=childChannelHandler
	return this
}

func (this *ServerBootstrap)Sync() *ServerBootstrap {

	_swait.Add(1)

	_swait.Wait()

	return this
}
