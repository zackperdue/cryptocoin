package main

import (
	"flag"

	"github.com/zackperdue/cryptocoin/network"
)

func main() {
	var port int
	flag.IntVar(&port, "p", 7203, "Blockchain Network Port Number")
	flag.Parse()

	network := network.Client{
		Port: port,
	}

	cryptocoin := Blockchain{
		Network: network,
	}

	cryptocoin.Init()
}
