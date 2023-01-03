package main

import (
	"backend/config"
	"backend/server"
	"os"
	"strconv"
	"strings"
)

func main() {

	// the only thing i know is what i have to do
	// and the servers in the network
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	hostname = strings.Split(hostname, ".")[0]
	N, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}

	all_servers := []string{}
	for i := 0; i < N; i++ {
		p := "node" + strconv.Itoa(i) + ":"
		all_servers = append(all_servers, p+config.DEFAULT_PORT)
	}
	server.Create(hostname, all_servers)

}
