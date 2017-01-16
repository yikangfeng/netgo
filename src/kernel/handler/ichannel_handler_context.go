package handler

type IChannelHandlerContext interface {
    Connect(host string,port int)
    Handler() IChannelHandler
}
