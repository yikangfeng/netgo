package handler

import (
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

func (this *TailContext) ChannelActive_(ctx common.IChannelHandlerContext) {
	//in netty action for set op_accept event
	fmt.Println("channel active")
}

func (this *TailContext) ChannelInactive_(ctx common.IChannelHandlerContext) {
	fmt.Println("channel inactive")
}
func (this *TailContext) ChannelRead_(ctx common.IChannelHandlerContext, msg interface{}) {

}

func (this *TailContext) ExceptionCaught(ctx common.IChannelHandlerContext, err error) {
	ctx.FireExceptionCaught(err)
}

