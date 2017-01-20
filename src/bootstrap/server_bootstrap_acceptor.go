package bootstrap

import (
	"kernel/intf/external/handler"
	"kernel/channel"
	handler_ "kernel/handler"
	"kernel/intf/external/common"
	"log"
	"time"
)

type ServerBootstrapAcceptor struct {
	//impl IChannelInboundHandler
	handler.IChannelInboundHandler
	ChildOption  map[string]interface{}
	ChildHandler common.IChannelHandler
}

func (this *ServerBootstrapAcceptor) ChannelActive_(ctx handler.IChannelHandlerContext) {
	//do nothing.
}

func (this *ServerBootstrapAcceptor) ChannelInactive_(ctx handler.IChannelHandlerContext) {
	//do nothing.
}

func (this *ServerBootstrapAcceptor) ChannelRead_(ctx handler.IChannelHandlerContext, msg interface{}) {
	if msg == nil {
		return
	}
	if (this.ChildHandler == nil) {
		return
	}
	socketChannel := msg.(*channel.SocketChannel)
	socketChannel.Config(this.ChildOption)
	if (this.ChildHandler != nil) {
		this.ChildHandler.(*handler_.ChannelInitializerHandler).ChannelInitFunc(socketChannel)
	}
	log.Println("start channel read...")
	socketChannel.Pipeline().FireChannelActive()
	go func() {
		//worker go
		var buffer []byte = make([]byte, 2048)
		for {
			len, err := socketChannel.Conn.Read(buffer)
			if err != nil {
				time.Sleep(1 * time.Second)
				continue
			}
                        log.Println(len)
			log.Println(string(buffer))

		}

	}()

}