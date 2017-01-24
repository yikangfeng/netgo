package protocols

import (
	"kernel/handler"
	"bytes"
	"kernel/intf/external/common"
	"encoding/binary"
	"errors"
)

type LengthFieldBasedFrameDecoder struct {
				//impl TransportProtocol
	handler.ChannelInboundHandlerAdapter
	packet              *bytes.Buffer
	MaxFrameLength      int
	LengthFieldOffset   int
	LengthFieldLength   int
	LengthAdjustment    int
	InitialBytesToStrip int /*表示在decode时，要去掉多少个字节 */
}

func NewLengthFieldBasedFrameDecoder(maxFrameLength int, lengthFieldOffset int, lengthFieldLength int, lengthAdjustment int, initialBytesToStrip int) (common.IChannelHandler) {
	return &LengthFieldBasedFrameDecoder{packet:bytes.NewBuffer(make([]byte, 0)), MaxFrameLength:maxFrameLength, LengthFieldOffset:lengthFieldOffset, LengthFieldLength:lengthFieldLength, LengthAdjustment:lengthAdjustment, InitialBytesToStrip:initialBytesToStrip}
}

func (this *LengthFieldBasedFrameDecoder)ChannelRead_(ctx common.IChannelHandlerContext, msg interface{}) {
	if _, ok := msg.(*bytes.Buffer); !ok {
		return
	}
	var _packet *bytes.Buffer = msg.(*bytes.Buffer)
	if (_packet == nil) {
		return
	}
	ibinaryMessage := NewBinaryMessage()
	if err, error := this.Unpacket(_packet.Bytes(), ibinaryMessage); (err == 0&&error == nil) {
		//read full packet.
		this.packet.Reset()
		go func() {
			ctx.FireChannelRead(ibinaryMessage)
		}()
	}

	this.packet.Write(_packet.Bytes())
}

func (this *LengthFieldBasedFrameDecoder)Packet(packet []byte, out IBinaryMessage) (int, error) {
	return 0, nil
}

func (this *LengthFieldBasedFrameDecoder)Unpacket(packet []byte, out IBinaryMessage) (int, error) {
	var lengthOffset int = this.LengthFieldOffset + this.LengthFieldLength
	var packetBytesLen = len(packet)
	if (packetBytesLen < lengthOffset) {
		return -1, errors.New("Unpack exception, no read to complete packet body.")
	}

	lengthBytes := packet[this.LengthFieldOffset:lengthOffset]

	if (len(lengthBytes) != this.LengthFieldLength) {
		return -1, errors.New("Unable to get the correct packet length.")
	}
	packetLength := int(binary.BigEndian.Uint32(lengthBytes))//packet body length.


	if (packetBytesLen != packetLength) {
		return -1, errors.New("Unpack exception, no read to complete packet body.")
	}

	var headerLengthOffset int = lengthOffset + 4
	headerLengthBytes := packet[lengthOffset:headerLengthOffset]
	headerLength := int(binary.BigEndian.Uint32(headerLengthBytes))

	var headerContentOffset int = headerLengthOffset + headerLength
	headerContentBytes := packet[headerLengthOffset:headerContentOffset]

	var packetBodyOffset int = (8 + headerLength)
	packetBodyBytes := packet[packetBodyOffset:]

	out.SetProtocolName(string(headerContentBytes))
	out.SetPacketHeader(&headerContentBytes)
	out.SetPacketBody(&packetBodyBytes)

	return 0, nil
}



