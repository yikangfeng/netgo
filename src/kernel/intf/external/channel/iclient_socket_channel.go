package channel

import "net"

type IClientSocketChannel interface {//impl ISocketChannel
	ISocketChannel
	ConnectAndInit(host string,port int)
	GetConn() (*net.TCPConn)
}