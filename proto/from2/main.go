package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	hello_world "proto/proto"
	pb "proto/proto"
)




func main(){
	service := micro.NewService(micro.Name("hello.client"))//客户端名称
	service.Init()
	// Create new client
	base := pb.NewHelloWorldService("hello_world", service.Client())
	res, err := base.Hello(context.TODO(), &hello_world.HelloRequest{Name:"wjw"})
	fmt.Println("res",res)
	fmt.Println("err",err)
}
