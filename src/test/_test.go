package test

import (
	"sync"
	"fmt"
	"os"
	"bootstrap"
)

func main() {

   _bootstrap:=bootstrap.New();
	_bootstrap.Channel(bootstrap.GetTCPChannel()).Bind(1024).Sync()
}

