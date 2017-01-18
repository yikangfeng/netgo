package channel

import (
	"net"
	"fmt"
	"os"
	"kernel/protocols"
	"time"
	"kernel/intf/external/channel"
)

type TCPSocketChannel struct {//impl IClientSocketChannel
	AbstractSocketChannel
}

func  NewTCPSocketChannel() (channel.IClientSocketChannel){
	instance:=&TCPSocketChannel{}
	instance.pipeline=NewChannelPipeline(instance)
	instance.config=make(map[string]interface{})
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
	config:=this.GetConfig();
	if(config==nil) {
		return
	}
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