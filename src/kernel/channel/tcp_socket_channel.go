package channel

import (
	"net"
	"fmt"
	"os"
	"kernel/protocols"
	"time"
)

type IClientSocketChannel interface {
	connectAndInit(host string,port int)
}
type TCPSocketChannel struct {//impl IClientSocketChannel
	AbstractSocketChannel
}

func  New() (*TCPSocketChannel){
	instance:=&TCPSocketChannel{}
	instance.pipeline=&ChannelPipeline{}
	instance.config=nil
      return instance
}


func (this *TCPSocketChannel) ConnectAndInit(host string,port int) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", fmt.Sprintf("%s:%d",host,port))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	conn,err := net.DialTCP(protocols.GetTCPProtocol().Name, nil, tcpAddr)
	if err != nil {
		defer conn.Close()
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	fmt.Println("connect success")
	this.init(conn)


}

func (this *TCPSocketChannel) init(conn *net.TCPConn) {
	if _,ok := this.config[ChannelOptions.Deadline];ok {
	   conn.SetDeadline(this.config[ChannelOptions.Deadline].(time.Time))
	}

	if _,ok := this.config[ChannelOptions.KeepAlive];ok {
		conn.SetKeepAlive(this.config[ChannelOptions.KeepAlive].(bool))
	}

	if _,ok := this.config[ChannelOptions.KeepAlivePeriod];ok {
		conn.SetKeepAlivePeriod(this.config[ChannelOptions.KeepAlivePeriod].(time.Duration))
	}

	if _,ok := this.config[ChannelOptions.Linger];ok {
		conn.SetLinger(this.config[ChannelOptions.Linger].(int))
	}
	if _,ok := this.config[ChannelOptions.NoDelay];ok {
		conn.SetNoDelay(this.config[ChannelOptions.NoDelay].(bool))
	}
	if _,ok := this.config[ChannelOptions.ReadBuffer];ok {
		conn.SetReadBuffer(this.config[ChannelOptions.ReadBuffer].(int))
	}
	if _,ok := this.config[ChannelOptions.ReadDeadline];ok {
		conn.SetReadDeadline(this.config[ChannelOptions.ReadDeadline].(time.Time))
	}
	if _,ok := this.config[ChannelOptions.WriteBuffer];ok {
		conn.SetWriteBuffer(this.config[ChannelOptions.WriteBuffer].(int))
	}
	if _,ok := this.config[ChannelOptions.WriteDeadline];ok {
		conn.SetWriteDeadline(this.config[ChannelOptions.WriteDeadline].(time.Time))
	}
}