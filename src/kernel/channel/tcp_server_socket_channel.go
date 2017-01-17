package channel

import (
	"kernel/protocols"
	"net"
	"fmt"
	"log"
	"os"
	"kernel/intf/external/channel"

)
type IServerSocketChannel interface {
	channel.ISocketChannel
	doBindAndAccept(host string,port int)
}
type TCPServerSocketChannel struct {
	AbstractSocketChannel
}


func  NewTCPServerSocketChannel() (IServerSocketChannel){
	instance:=&TCPServerSocketChannel{}
	instance.pipeline=NewChannelPipeline(instance)
	instance.config=make(map[string]interface{})
	return instance
}

func (this *TCPServerSocketChannel)doBindAndAccept(host string,port int)  {
	netListener, err := net.Listen(protocols.GetTCPProtocol().Name, fmt.Sprintf("%s:%d",host,port))
	checkError(err)
	defer netListener.Close()
	Log("Waiting for clients")

	go func() {
		for {
			Log("start blocking...")
			conn, err := netListener.Accept()
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
func Log(v ...interface{}) {
	log.Println(v)
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}