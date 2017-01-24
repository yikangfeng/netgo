package string

import (
	"kernel/handler"
)
/*
包长+包头标记+协议名称+包体
len+<BOF>+协议名称+包体
 */

type StringPacket struct {
	//impl TransportMessageProtocol extends ChannelOutboundHandlerAdapter
	handler.ChannelOutboundHandlerAdapter
}

func (this *StringPacket)Packet(msg interface{}, out interface{}) (int, error) {
	return 0, nil
}

func (this *StringPacket)Unpacket(msg interface{}, out interface{}) (int, error) {
	return 0, nil
}