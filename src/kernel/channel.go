package kernel

import (
	"sync"
)

type Channel struct {
      Name string
}

var tcpChannel,udpChannel *Channel


var _lock sync.Mutex
func GetTCPChannel() *Channel{

	_lock.Lock()
	if(tcpChannel==nil){
		tcpChannel=&Channel{Name:"tcp"}
	}
         _lock.Unlock()
	return tcpChannel
}

func GetUDPChannel()  *Channel{
	_lock.Lock()
	if(tcpChannel==nil){
		udpChannel=&Channel{Name:"udp"}
	}
	_lock.Lock()
	return udpChannel
}