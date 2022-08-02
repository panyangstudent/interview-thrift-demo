package main

import (
	"flag"
	"interview/gen-go/user"
)

func main() {
	User()
}

func User() {
	server := flag.Bool("server",true,"是否是服务器")
	flag.Parse()

	if *server {
		user.UserService()
	} else {
		user.UserServiceClient{}
		simple_client.SimpleClient()
	}
}
