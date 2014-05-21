package data

import (
	"github.com/proj-223/CatFs/config"
	proc "github.com/proj-223/CatFs/protocols"
	"github.com/proj-223/CatFs/utils"
)

const (
	RPC_START_MSG = "CatFS Data Server %d RPC are start: %s\n"
)

var (
	DefaultDataServers []*DataServer = []*DataServer{
		NewDataServer(config.DefaultMachineConfig, 0),
		NewDataServer(config.DefaultMachineConfig, 1),
		NewDataServer(config.DefaultMachineConfig, 2),
		NewDataServer(config.DefaultMachineConfig, 3),
	}
)

func Serve(index int) error {
	return ServeDataServer(DefaultDataServers[index])
}

func ServeDataServer(data *DataServer) error {
	done := make(chan error, 1)

	// init the rpc server
	go data.initRPCServer(done)
	// init the block server
	go data.initBlockServer(done)

	err := <-done
	return err
}

// Create a new Master Server
func NewDataServer(conf *config.MachineConfig, index int) *DataServer {
	ds := &DataServer{
		pool:        proc.NewClientPool(conf),
		conf:        conf,
		index:       index,
		blockServer: utils.NewBlockServer(conf.BlockServerConf),
	}
	return ds
}