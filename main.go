package main

import (
	"flag"
	"interview/client"
	"interview/server"
)

func main() {
	Simple()
}

func Simple() {
	server1 := flag.Bool("server",true,"是否是服务器")
	flag.Parse()

	if *server1 {
		server.SimpleServer()
	} else {
		client.SimpleClient()
	}
}