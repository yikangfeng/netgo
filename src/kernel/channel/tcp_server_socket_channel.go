package channel

import (
	"kernel/protocols"
	"net"
	"fmt"
	"log"
	"os"
	"kernel/intf/external/channel"
	"time"
)

type TCPServerSocketChannel struct {
	//impl IServerSocketChannel
	AbstractSocketChannel
}

func NewTCPServerSocketChannel() (channel.IServerSocketChannel) {
	instance := &TCPServerSocketChannel{}
	instance.pipeline = NewChannelPipeline(instance)
	instance.config = make(map[string]interface{})
	instance.parent = nil
	return instance
}

var netListener net.Listener
var err error

func (this *TCPServerSocketChannel)DoClose() {
	defer netListener.Close()
}
func (this *TCPServerSocketChannel)DoBindAndAccept(host string, port int) {
	netListener, err = net.Listen(protocols.GetTCPProtocol().Name, fmt.Sprintf("%s:%d", host, port))
	checkError(err)
	Log("bind success.")
	this.doAccept()
}

func (this *TCPServerSocketChannel)doAccept() {
	Log("start accept...")
	go func() {
		for {
			conn, err := netListener.Accept()
			if err != nil {
				//not available conn.
				Log("not available conn sleep 1s...")
				time.Sleep(1 * time.Second)
			}
			Log(conn.RemoteAddr().String(), " client tcp connect accepted success.")
			this.Pipeline().FireChannelRead(NewSocketChannel(this, conn))
		}

	}()
}

func Log(v ...interface{}) {
	log.Println(v)
}
func checkError(err error) {
	if err != nil {
		defer netListener.Close()
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}