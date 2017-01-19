package channel

import (
	"kernel/intf/external/channel"
	"kernel/intf/external/common"
	"net"
)

type SocketChannel struct {
	AbstractSocketChannel
	Conn net.Conn
}

func NewSocketChannel(_parent common.IChannel,_conn net.Conn) (channel.ISocketChannel) {
	instance := &SocketChannel{}
	instance.pipeline = NewChannelPipeline(instance)
	instance.Conn = _conn
	instance.parent=_parent
	return instance
}