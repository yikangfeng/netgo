package channel


type IServerSocketChannel interface {//impl ISocketChannel
	ISocketChannel
	DoBindAndAccept(host string,port int)
}