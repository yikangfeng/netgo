package protocols

type IBinaryMessage interface {
	GetProtocolName() (string)
	SetProtocolName(protocolName string)
	GetPacketHeader() (*[]byte)
	SetPacketHeader(packetHeader *[]byte)
	GetPacketBody() (*[]byte)
	SetPacketBody(packetBody *[]byte)
}

type BinaryMessage struct {
	//impl IMessage
	ProtocolName string
	PacketHeader *[]byte
	PacketBody   *[]byte
}

func NewBinaryMessage() IBinaryMessage {
	return &BinaryMessage{ProtocolName:"", PacketHeader:nil, PacketBody:nil}
}

func (this *BinaryMessage) GetProtocolName() (string) {
	return this.ProtocolName
}

func (this *BinaryMessage) SetProtocolName(protocolName string) {
	this.ProtocolName = protocolName
}

func (this *BinaryMessage) GetPacketHeader() (*[]byte) {
	return this.PacketHeader
}

func (this *BinaryMessage) SetPacketHeader(packetHeader *[]byte) {
	this.PacketHeader = packetHeader
}

func (this *BinaryMessage) GetPacketBody() (*[]byte) {
	return this.PacketBody
}

func (this *BinaryMessage) SetPacketBody(packetBody *[]byte) {
	this.PacketBody = packetBody
}