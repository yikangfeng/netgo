package handler

import (
	"kernel/intf/external/handler"
	//"kernel/intf/external/channel"

)

type TailContext struct {
					      //impl IChannelInboundHandler
	AbstractChannelHandlerContext
}

func (this *TailContext) Bind(host string, port int) {

}

func (this *TailContext) Connect(host string, port int) {

}

func (this *TailContext) Handler() (handler.IChannelHandler) {
	return this
}
