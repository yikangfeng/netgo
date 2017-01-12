package client

import (
"sync"
	"kernel"
)

/*
@author YiKangfeng.
 */
type ClientBootstrap struct {

	option map[string]interface{}

	_channel *kernel.Channel


}

var _wait sync.WaitGroup

func  New() *ClientBootstrap{
	return &ClientBootstrap{option:make(map[string]interface{})}
}

func (this *ClientBootstrap)Channel( channel *kernel.Channel) (_this *ClientBootstrap) {
	if(channel==nil){
		return nil
	}
	this._channel=channel;
	return this
}

func (this *ClientBootstrap)Option( key  string, v interface{}) *ClientBootstrap {
	_,ok := this.option[key]
	if(!ok) {
		this.option[key] = v
	}
	return this
}



func (this *ClientBootstrap)Handler( key  string, v interface{}) *ClientBootstrap {
	return this
}

func (this *ClientBootstrap)Connect(port int) *ClientBootstrap {

	return this
}
func (this *ClientBootstrap)Sync() *ClientBootstrap {
	return this
}