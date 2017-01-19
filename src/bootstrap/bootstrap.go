package bootstrap

import (
	"sync"
	"kernel/intf/external/common"
)

/*
@author YiKangfeng.
 */
type Bootstrap struct {
	//extends abstractbootstrap

	abstractbootstrap
}

var _wait sync.WaitGroup

func NewBootstrap() *Bootstrap {
	bootstrap := &Bootstrap{}
	bootstrap.option = make(map[string]interface{})
	return bootstrap
}

func (this *Bootstrap)Option(key  string, v interface{}) *Bootstrap {
	_, ok := this.option[key]
	if (!ok) {
		this.option[key] = v
	}
	return this
}

func (this *Bootstrap)Channel(_channel common.IChannel) *Bootstrap {
	if (_channel == nil) {
		return nil
	}
	this.channel = _channel;
	return this
}

//for Logging Handler
func (this *Bootstrap)Handler(channelHandler common.IChannelHandler) *Bootstrap {
	this.handler = channelHandler
	return this
}

func (this *Bootstrap)Connect(host string, port int) *Bootstrap {
	this.init()
	go func() {
		this.channel.Connect(host, port)
	}()
	return this
}

func (this *Bootstrap)Sync() *Bootstrap {
	_wait.Add(1)
	_wait.Wait()
	return this
}