package protocols

type ITransportMessageProtocol interface {
	Packet(msg interface{}) (int, error, interface{})
	Unpacket(msg interface{}) (int, error, interface{})
}