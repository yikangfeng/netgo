package string

import (
	"kernel/handler"
	"kernel/intf/external/common"
	"fmt"
	"kernel/handler/transport/protocols"
)
/*
包长+包头长度+包头内容+包体
len+headerlength+包头内容+包体
 */

type StringUnpacket struct {
	//impl TransportMessageProtocol extends ChannelInboundHandlerAdapter
	handler.ChannelInboundHandlerAdapter
}

func NewStringUnpacket() (common.IChannelHandler) {
	return &StringUnpacket{}
}

func (this *StringUnpacket)Packet(msg interface{}) (int, error, interface{}) {
	return 0, nil, nil
}

func (this *StringUnpacket)Unpacket(msg interface{}) (int, error, interface{}) {

	return 0, nil, string(*msg.(protocols.IBinaryMessage).GetPacketBody())
}

func (this *StringUnpacket)ChannelRead_(ctx common.IChannelHandlerContext, msg interface{}) {

	if binaryMessage, ok := msg.(protocols.IBinaryMessage); ok {
		_, _, content := this.Unpacket(binaryMessage)
		fmt.Println(content)
		ctx.FireChannelRead(content)
	}

}