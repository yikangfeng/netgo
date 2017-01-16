package channel



type AbstractChannel struct {//impl IChannel
	IChannel
}


type AbstractSocketChannel struct {//extends AbstractChannel
	AbstractChannel
	config *map[string]interface{}
	pipeline *ChannelPipeline
}


func (this *AbstractSocketChannel)Pipeline() *ChannelPipeline{
  return this.pipeline
}

func (this *AbstractSocketChannel)Close(){

}

func (this *AbstractSocketChannel)Config(_config *map[string]interface{}){
     this.config=_config
}

func (this *AbstractSocketChannel)GetConfig() *map[string]interface{}{
	return this.config
}


func (this *AbstractSocketChannel)Connect(host string,port int)  {
	this.pipeline().Connect(host,port)
}