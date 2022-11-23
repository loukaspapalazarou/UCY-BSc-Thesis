// Client

package main

import (
	"fmt"
	"frontend/client"
	"frontend/config"
	"frontend/messaging"
	"time"
)

func client_task(id string, servers []config.Node) {

	client := client.Create(id, servers)

	// messaging.Add(client, "Hello")
	messaging.TargetedAdd(client, *client.Servers[0], "Hello")
	messaging.TargetedAdd(client, *client.Servers[1], "World")
	messaging.TargetedAdd(client, *client.Servers[2], "!!!!")
	time.Sleep(time.Second * 10)
	g, _ := messaging.GetGset(client)
	fmt.Println(g)
	// for {
	// 	messaging.TargetedAdd(client, *client.Servers[0], "Hello")
	// 	time.Sleep(time.Second * 2)
	// }

	// time.Sleep(time.Second * 3)
	// messaging.TargetedAdd(client, *client.Servers[0], "Hello")
	// messaging.TargetedAdd(client, *client.Servers[1], "Hello22")
	// s, _ := messaging.GetGset(client)
	// fmt.Println(s)
	// messaging.SimpleBroadcast([]string{messaging.GET}, server_sockets)

}

func main() {

	LOCAL := true
	var servers []config.Node
	if LOCAL {
		servers = config.Servers_LOCAL
	} else {
		servers = config.Servers
	}

	go client_task("c1", servers)

	for {
	}

}
