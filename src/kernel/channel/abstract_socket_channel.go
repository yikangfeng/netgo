package channel

import (
	"kernel/intf/external/channel"
	"kernel/intf/external/common"
)

type AbstractChannel struct {
	//impl IChannel
	channel.ISocketChannel
	pipeline common.IChannelPipeline
	parent common.IChannel
}

type AbstractSocketChannel struct {
	//extends AbstractChannel
	AbstractChannel
	config map[string]interface{}
}

func (this *AbstractChannel)Close() {

}
func (this *AbstractChannel)Pipeline() (common.IChannelPipeline) {
	return this.pipeline
}

func (this *AbstractChannel)Connect(host string, port int) {
	this.Pipeline().Connect(host, port)
}

func (this *AbstractChannel)Bind(host string, port int) {
	this.Pipeline().Bind(host, port)
}

func (this *AbstractSocketChannel)Config(config map[string]interface{}) {
	this.config = config
}

func (this *AbstractSocketChannel)GetConfig() map[string]interface{} {
	return this.config
}

func (this *AbstractSocketChannel)GetParent() (common.IChannel) {
	return this.parent
}