package main

import (
	"context"
	"fmt"
	client "github.com/micro/go-micro/client"
	hello_world "proto/proto"
)

func main(){
	//调用end1
	//client.Client类型
	var c  = client.NewClient()
	name := "hello_world"
	//Base.User
	fmt.Println("1ssss")
	req := c.NewRequest(name, "HelloWorld.Hello",&hello_world.HelloRequest{Name:"wjw"})

	out := new(hello_world.HelloResponse)
	err := c.Call(context.TODO(), req, out)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(out)

}