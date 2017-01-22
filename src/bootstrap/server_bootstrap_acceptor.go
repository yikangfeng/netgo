package bootstrap

import (
	"kernel/channel"
	handler_ "kernel/handler"
	"kernel/intf/external/common"
	"log"
	"net"
)

type ServerBootstrapAcceptor struct {
	//impl IChannelInboundHandler
	handler_.ChannelHandlerAdapter
	ChildOption  map[string]interface{}
	ChildHandler common.IChannelHandler
}

func (this *ServerBootstrapAcceptor) ChannelActive_(ctx common.IChannelHandlerContext) {
	//do nothing.
}

func (this *ServerBootstrapAcceptor) ChannelInactive_(ctx common.IChannelHandlerContext) {
	//do nothing.
}
func (this *ServerBootstrapAcceptor) ExceptionCaught_(ctx common.IChannelHandlerContext, err error) {
	//do nothing.
}

func (this *ServerBootstrapAcceptor) ChannelRead_(ctx common.IChannelHandlerContext, msg interface{}) {
	if msg == nil {
		return
	}
	if (this.ChildHandler == nil) {
		return
	}
	socketChannel := msg.(*channel.SocketChannel)
	socketChannel.Config(this.ChildOption)
	if (this.ChildHandler != nil) {
		socketChannel.Pipeline().AddLast(this.ChildHandler)
		this.ChildHandler.(*handler_.ChannelInitializerHandler).ChannelInitFunc(socketChannel)
	}
	log.Println("start channel read...")
	socketChannel.Pipeline().FireChannelActive()
	go func() {
		//worker go
		var cchannel net.Conn = socketChannel.Conn;
		defer cchannel.Close()
		var buffer []byte = make([]byte, 2048)
		for {
			len, err := cchannel.Read(buffer)
			if err != nil {
				//has error.
				socketChannel.Pipeline().FireExceptionCaught(err)
				break
			}
			log.Println(len)
			log.Println(string(buffer))

		}

	}()

}