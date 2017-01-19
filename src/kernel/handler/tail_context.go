package handler

import (
	"kernel/intf/external/handler"
	"kernel/intf/external/common"

	"fmt"
)

type TailContext struct {
	//impl IChannelInboundHandler
	AbstractChannelHandlerContext
}


func (this *TailContext) Handler() (common.IChannelHandler) {
	return this
}

func (this *TailContext) ChannelActive_(ctx handler.IChannelHandlerContext) {
	//in netty action for set op_accept event
	fmt.Println("channel active")
}

func (this *TailContext) ChannelInactive_(ctx handler.IChannelHandlerContext) {
	fmt.Println("channel inactive")
}
func (this *TailContext) ChannelRead_(ctx handler.IChannelHandlerContext, msg interface{}) {

}



