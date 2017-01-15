package socket

import (
	"kernel/channel"
	"net"
	"fmt"
	"os"
	"kernel/handler"
	"kernel/protocols"
)

type IClientSocketChannel interface {
	channel.IChannel
	Connect(host string,port int)
}
type TCPSocketChannel struct {

}

func (this *TCPSocketChannel) Connect(host string,port int) {
	go func() {
		tcpAddr, err := net.ResolveTCPAddr("tcp4", fmt.Sprintf("%s:%d",host,port))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
			os.Exit(1)
		}

		conn,err := net.DialTCP(protocols.GetTCPProtocol(), nil, tcpAddr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
			os.Exit(1)
		}
		fmt.Println("connect success")

		handler.NewChannelHandlerContext(conn)
	}()
}

func (this *TCPSocketChannel) Pipeline() (channelPipeline *channel.ChannelPipeline) {
      return nil
}