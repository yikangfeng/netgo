package channel


type IChannel interface {
     Pipeline() (channelPipeline *ChannelPipeline)
}