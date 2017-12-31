package network

import (
	"fmt"
	"log"
	"net"

	"github.com/go-redis/redis"
)

type Client struct {
	Port         int
	redis        *redis.Client
	networkPeers []string
}

func (c Client) Connect() {
	c.redis = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := c.redis.Ping().Result()

	if err != nil {
		log.Fatalln("Coudn't connect to server list")
	}

	fmt.Println(fmt.Sprintf("Connection to local server list established: %s.", pong))

	c.registerSelf()
	c.discoverPeers()
}

func (c Client) registerSelf() {
	c.redis.SAdd("nodes", fmt.Sprintf("%s:%d", c.getOutboutIP(), c.Port))
}

func (c Client) getOutboutIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatalln("My path to existance has failed.")
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}

	return ""
}

func (c Client) discoverPeers() {
	nodes, err := c.redis.SMembersMap("nodes").Result()

	if err != nil {
		log.Fatalln("Failed to get node list")
	}

	for address := range nodes {
		fmt.Println("Address: ", address)
		c.networkPeers = append(c.networkPeers, address)
	}
}
