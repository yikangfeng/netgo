package bootstrap

import (
	"sync"
	"kernel/channel"
	"kernel/handler"
	"kernel/channel/socket"
)

/*
@author YiKangfeng.
 */
type ServerBootstrap struct {

 option map[string]interface{}

 childOption map[string]interface{}

 channel channel.IChannel


}

var _wait sync.WaitGroup

func  New() *ServerBootstrap{
	return &ServerBootstrap{option:make(map[string]interface{}),
	childOption:make(map[string]interface{})}
}

func (this *ServerBootstrap)Channel( _channel channel.IChannel) (_this *ServerBootstrap) {
	if(_channel==nil){
		return nil
	}
	this.channel=_channel;
	return this
}

func (this *ServerBootstrap)Option( key  string, v interface{}) *ServerBootstrap {
	_,ok := this.option[key]
	if(!ok) {
		this.option[key] = v
	}
	return this
}

func (this *ServerBootstrap)ChildOption( key  string, v interface{}) *ServerBootstrap {
	_,ok := this.option[key]
	if(!ok) {
		this.option[key] = v
	}
        return this
}

func (this *ServerBootstrap)Handler( channelHandler handler.ChannelHandler) *ServerBootstrap {
	return this
}

func (this *ServerBootstrap)ChildHandler( channelHandler handler.ChannelHandler) *ServerBootstrap {
	return this
}

func (this *ServerBootstrap)Bind(port int) *ServerBootstrap {

	this.channel.(socket.IServerSocketChannel).DoBindAndAccept(port)

	return this
}
func (this *ServerBootstrap)Sync() *ServerBootstrap {

	_wait.Add(1)

	_wait.Wait()

	return this
}
