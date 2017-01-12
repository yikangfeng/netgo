package bootstrap

import (
	"net"
	"log"
	"fmt"
	"os"
	"sync"
)

/*
@author YiKangfeng.
 */
type ServerBootstrap struct {

 option map[string]interface{}

 childOption map[string]interface{}

 _channel *Channel


}

var _wait sync.WaitGroup

func  New() *ServerBootstrap{
	return &ServerBootstrap{option:make(map[string]interface{}),
	childOption:make(map[string]interface{})}
}

func (this *ServerBootstrap)Channel( channel *Channel) (_this *ServerBootstrap) {
	if(channel==nil){
		return nil
	}
	this._channel=channel;
	return this
}

func (this *ServerBootstrap)Option( key  string, v interface{}) *ServerBootstrap {
	_,ok := this.option[key]
	if(!ok) {
		this.option[key] = v
	}
	return this
}

func (this *ServerBootstrap)ChildOption( key  string, v interface{}) *ServerBootstrap {
	_,ok := this.option[key]
	if(!ok) {
		this.option[key] = v
	}
        return this
}


func (this *ServerBootstrap)ChildHandler( key  string, v interface{}) *ServerBootstrap {
	return this
}

func (this *ServerBootstrap)Bind(port int) *ServerBootstrap {


	go func(){
		netListen, err := net.Listen(this._channel.Name, fmt.Sprintf("%s:%d","127.0.0.1",port))
		checkError(err)
		defer netListen.Close()
		Log("Waiting for clients")
		for {
			Log("start blocking...")
			conn, err := netListen.Accept()
			Log("accepted conn.")
			if err != nil {
				continue
			}

			Log(conn.RemoteAddr().String(), " tcp connect success")

			go func() {
                            handleConnection(conn)
			}()

		}

	}()


	return this
}
func (this *ServerBootstrap)Sync() *ServerBootstrap {

	_wait.Add(1)

	_wait.Wait()

	return this
}

func handleConnection(conn net.Conn) {

	buffer := make([]byte, 2048)

	for {

		Log("reading...")
		n, err := conn.Read(buffer)
		Log("reading completed.")
		if err != nil {

			Log(conn.RemoteAddr().String(), " connection error: ", err)
			//_wait.Done()
			return
		}

		Log(conn.RemoteAddr().String(), "receive data string:\n", string(buffer[:n]))

	}

}
func Log(v ...interface{}) {
	log.Println(v)
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
