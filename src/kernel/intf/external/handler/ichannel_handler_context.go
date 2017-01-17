package handler

type IChannelHandlerContext interface {//handler context
    Connect(host string,port int)
    Bind(host string,port int)
    Handler() IChannelHandler
    Next() IChannelHandlerContext
    SetNext(ctx IChannelHandlerContext)
    Prev() IChannelHandlerContext
    SetPrev(ctx IChannelHandlerContext)
    GetInbound() bool
    GetOutbound() bool
}
