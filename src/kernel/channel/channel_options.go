package channel


type channel_options struct {
       Deadline string//int
       KeepAlive string//bool
       KeepAlivePeriod string//time.Duration
       Linger string//int
       NoDelay string//bool
       ReadBuffer string//int
       ReadDeadline string//time.Time
       WriteBuffer string//int
       WriteDeadline string//time.Time
}

var ChannelOptions *channel_options=&channel_options{"deadline","keep_alive","keep_alive_period","linger","no_delay","readbuffer","read_deadline","writebuffer","write_deadline"}




