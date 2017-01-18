package channel

type IClientSocketChannel interface {//impl ISocketChannel
	ISocketChannel
	ConnectAndInit(host string,port int)
}