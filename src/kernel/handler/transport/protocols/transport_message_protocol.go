package protocols

type ITransportMessageProtocol interface {
	Packet(msg interface{}, out interface{}) (int, error)
	Unpacket(msg interface{}, out interface{}) (int, error)
}