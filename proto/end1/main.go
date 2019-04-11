package main

import (
	"fmt"
	//hello_world "proto/modelsone"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"golang.org/x/net/context"
	hello_world"proto/proto"
)


type helloWorldClient struct {
	c           client.Client
	serviceName string
}

//执行命令 protoc --proto_path= D:\goprj\src\proto\models\hello_world.proto  --micro_out=.  --go_out=. -I D:\goprj\src\proto\

func main(){

	service := micro.NewService(
		micro.Name("hello_world"),
		micro.Version("latest"),
		micro.Metadata(map[string]string{
			"type": "helloworld",
		}),
	)
	// # 调用micro.NewService(调用micro.newService)来创建实现micro.Service interface的micro.service
	service.Init()//# 调用micro.service#Init方法
	hello_world.RegisterHelloWorldHandler(service.Server(), new(HelloWorld))
	service.Run() //# 调用micro.service#Run方法

}

type HelloWorld struct{}

func (g *HelloWorld) Hello(ctx context.Context, req *hello_world.HelloRequest, rsp *hello_world.HelloResponse) error {
	fmt.Println("Hello World",req.Name)
	//Id  HelloRequest返回中和对方都有Id同样的类型就能接收到
	rsp.Id = 123456
	fmt.Println("1sssz")
	return nil
} // 实现hello_world service中Hello方法





/*

调用端代码
func (c *baseService) User(ctx context.Context, in *UserRequest, opts ...client.CallOption) (*UserReply, error) {
	c.name = "hello_world"
	//Base.User
	fmt.Println("1ssss")
	req := c.c.NewRequest(c.name, "HelloWorld.Hello", in)
	out := new(UserReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

 */