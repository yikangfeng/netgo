package common

type IChannelHandler interface {
	ExceptionCaught(ctx IChannelHandlerContext, err error)
}