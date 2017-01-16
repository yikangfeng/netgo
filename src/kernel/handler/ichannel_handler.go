package handler



type IChannelHandler interface {
	HandlerAdded() (int,error)

	HandlerRemoved() (int,error)
}