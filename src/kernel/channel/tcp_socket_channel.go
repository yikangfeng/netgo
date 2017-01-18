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

func  NewTCPSocketChannel() (*TCPSocketChannel){
	instance:=&TCPSocketChannel{}
	instance.pipeline=&ChannelPipeline{}
	instance.config=nil
      return instance
}


func (this *TCPSocketChannel) connectAndInit(host string,port int) {
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
	config:=this.GetConfig();
	if _,ok := config[ChannelOptions.Deadline];!ok {
	   conn.SetDeadline(this.config[ChannelOptions.Deadline].(time.Time))
	}

	if _,ok := config[ChannelOptions.KeepAlive];!ok {
		conn.SetKeepAlive(this.config[ChannelOptions.KeepAlive].(bool))
	}

	if _,ok := config[ChannelOptions.KeepAlivePeriod];!ok {
		conn.SetKeepAlivePeriod(this.config[ChannelOptions.KeepAlivePeriod].(time.Duration))
	}

	if _,ok := config[ChannelOptions.Linger];!ok {
		conn.SetLinger(this.config[ChannelOptions.Linger].(int))
	}
	if _,ok := config[ChannelOptions.NoDelay];!ok {
		conn.SetNoDelay(this.config[ChannelOptions.NoDelay].(bool))
	}
	if _,ok := config[ChannelOptions.ReadBuffer];!ok {
		conn.SetReadBuffer(this.config[ChannelOptions.ReadBuffer].(int))
	}
	if _,ok := config[ChannelOptions.ReadDeadline];!ok {
		conn.SetReadDeadline(this.config[ChannelOptions.ReadDeadline].(time.Time))
	}
	if _,ok := config[ChannelOptions.WriteBuffer];!ok {
		conn.SetWriteBuffer(this.config[ChannelOptions.WriteBuffer].(int))
	}
	if _,ok := config[ChannelOptions.WriteDeadline];!ok {
		conn.SetWriteDeadline(this.config[ChannelOptions.WriteDeadline].(time.Time))
	}
}