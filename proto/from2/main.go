package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	hello_world "proto/proto"
	pb "proto/proto"
)




func main(){
	//hello.client这里填写结构的名称 hello为结构 client为方法或者直接写client也行
	service := micro.NewService(micro.Name("hello.client"))//客户端名称
	service.Init()
	// Create new client hello_world为服务名称
	base := pb.NewHelloWorldService("hello_world", service.Client())
	res, err := base.Hello(context.TODO(), &hello_world.HelloRequest{Name:"wjw"})
	fmt.Println("res",res)
	fmt.Println("err",err)
}
