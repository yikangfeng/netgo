package channel


type IChannel interface {
     Pipeline() (channelPipeline *ChannelPipeline)
     Close()
     Config(config map[string]interface{})
     Connect(host string,port int)
     Bind(host string,port int)
}