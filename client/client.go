package client

import (
	"context"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"interview/gen-go/Sample"
	"time"
)

func handleClient(client *Sample.SimpleServiceClient) {
	ctx := context.TODO()
	res, _ := client.Add(ctx,13,"25")
	fmt.Println("result is ", res)
}
func SimpleClient()  {
	var transport thrift.TTransport
	transport,  _ = thrift.NewTSocketConf("localhost:8090", &thrift.TConfiguration{
		ConnectTimeout: time.Second, // Use 0 for no timeout
		SocketTimeout:  time.Second, // Use 0 for no timeout
	})
	conf := &thrift.TConfiguration{
		MaxFrameSize:       1024*256,
		ConnectTimeout:     time.Second,
		SocketTimeout:      time.Second,
		TBinaryStrictRead:  thrift.BoolPtr(true),
		TBinaryStrictWrite: thrift.BoolPtr(true),
	}
	protocolFactory := thrift.NewTBinaryProtocolFactoryConf(conf)
	transportFactory := thrift.NewTTransportFactory()
	transport, _ = transportFactory.GetTransport(transport)
	defer transport.Close()
	transport.Open()
	iprot := protocolFactory.GetProtocol(transport)
	oprot := protocolFactory.GetProtocol(transport)
	handleClient(Sample.NewSimpleServiceClient(thrift.NewTStandardClient(iprot, oprot)))
}