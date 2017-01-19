package common

type IChannel interface {
	Pipeline() (IChannelPipeline)
	Close()
	Config(config map[string]interface{})
	GetConfig() (map[string]interface{})
	GetParent() (IChannel)
	Connect(host string, port int)
	Bind(host string, port int)
}