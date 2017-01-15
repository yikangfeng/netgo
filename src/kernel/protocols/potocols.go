package protocols


type Protocol struct {
      Name string
}

var tcp *Protocol=&Protocol{Name:"tcp"}
var udp *Protocol=&Protocol{Name:"udp"}



func GetTCPProtocol() *Protocol{
	return tcp
}

func GetUDPProtocol()  *Protocol{
	return udp
}