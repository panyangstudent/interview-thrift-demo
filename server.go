package interview_thrift_demo

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"interview/gen-go/user"
	"time"
)

func UserServer()  {
	conf := &thrift.TConfiguration{
		ConnectTimeout: time.Second,
		SocketTimeout:  time.Second,
		MaxFrameSize: 1024 * 256,
		TBinaryStrictRead:  thrift.BoolPtr(true),
		TBinaryStrictWrite: thrift.BoolPtr(true),
	}
	protocolFactory := thrift.NewTBinaryProtocolFactoryConf(conf)
	transportFactory :=  thrift.NewTTransportFactory()

	transport, _:= thrift.NewTServerSocket("8090")
	processor := user.NewUserServiceProcessor(&UserServiceHandler{})
	server := thrift.NewTSimpleServer4(processor,transport,transportFactory,protocolFactory)
	server.Serve()
}