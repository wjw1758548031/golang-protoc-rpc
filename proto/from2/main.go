package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	hello_world "proto/proto"
	pb "proto/proto"
)




func main(){
	//hello.client 要么是结构名，要么是rpc里通道定义的名称
	service := micro.NewService(micro.Name("hello.client"))//客户端名称
	service.Init()
	// Create new client hello_world为服务名称
	base := pb.NewHelloWorldService("hello_world", service.Client())
	res, err := base.Hello(context.TODO(), &hello_world.HelloRequest{Name:"wjw"})
	fmt.Println("res",res)
	fmt.Println("err",err)
}
