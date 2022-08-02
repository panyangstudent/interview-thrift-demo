package interview_thrift_demo

import (
	"flag"
	"interview/server"
	"interview/client"
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