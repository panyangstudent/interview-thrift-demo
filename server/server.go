package server

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"interview/handler"
	"interview/gen-go/Sample"
	"time"
)

func SimpleServer() {
	conf := &thrift.TConfiguration{
		ConnectTimeout: time.Second,
		SocketTimeout:  time.Second,
		MaxFrameSize: 1024 * 256,
		TBinaryStrictRead:  thrift.BoolPtr(true),
		TBinaryStrictWrite: thrift.BoolPtr(true),
	}

	protocolFactory := thrift.NewTBinaryProtocolFactoryConf(conf)
	transportFactory := thrift.NewTTransportFactory()

	//
	transport, _ := thrift.NewTServerSocket(":8090")

	processor := Sample.NewSimpleServiceProcessor(&handler.SimpleServiceHandler{})
	//阻塞式单线程服务器，阻塞式IO
	server := thrift.NewTSimpleServer4(processor,transport,transportFactory,protocolFactory)
	server.Serve()
}