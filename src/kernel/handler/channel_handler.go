package handler



type ChannelHandler interface {
	HandlerAdded() (int,error)

	HandlerRemoved() (int,error)

}