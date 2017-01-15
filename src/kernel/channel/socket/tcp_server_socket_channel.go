package socket

import (
	"kernel/protocols"
	"net"
	"fmt"
	"log"
	"os"
	"kernel/channel"
)
type IServerSocketChannel interface {
	channel.ISocketChannel
	DoBindAndAccept(port int)
}
type TCPServerSocketChannel struct {

}

func (this *TCPServerSocketChannel)DoBindAndAccept(port int)  {

	netListen, err := net.Listen(protocols.GetTCPProtocol().Name, fmt.Sprintf(":%d",port))
	checkError(err)
	defer netListen.Close()
	Log("Waiting for clients")

	go func() {
		for {
			Log("start blocking...")
			conn, err := netListen.Accept()
			Log("accepted conn.")
			if err != nil {
				continue
			}

			Log(conn.RemoteAddr().String(), " tcp connect success")

			go func() {
				//handle (conn)
			}()

		}
	}()



}

func (this *TCPServerSocketChannel) Pipeline() (channelPipeline *channel.ChannelPipeline) {
	return  nil
}

func Log(v ...interface{}) {
	log.Println(v)
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}