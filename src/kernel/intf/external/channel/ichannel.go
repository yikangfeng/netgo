package channel


type IChannel interface {
     Pipeline() (IChannelPipeline)
     Close()
     Config(config map[string]interface{})
     Connect(host string,port int)
     Bind(host string,port int)
}