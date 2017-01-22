package codec

type Codec interface {
	Encode()
	Decode()
}