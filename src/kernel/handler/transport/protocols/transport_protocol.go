package protocols

type ITransportProtocol interface {
	Packet(msg []byte, out IBinaryMessage) (int, error)
	Unpacket(msg []byte, out IBinaryMessage) (int, error)
}