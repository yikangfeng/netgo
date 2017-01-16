package handler


type IChannelOutboundHandler interface {
	IChannelHandler
	Connect(host string,port int)
}


