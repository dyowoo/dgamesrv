/**
* @Author: Jason
* @Date: 2020/9/30 9:46
* @File : main
* @Software: GoLand
**/

package main

import (
	"dygame/dcore"
	protoMsg "dygame/message"
	"log"
	"math/rand"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var con = dcore.SqlCon{}

type server struct{}

func initRpc() {
	listen, err := net.Listen("tcp", ":18029")
	if err != nil {
		log.Fatalf("监听失败：%v", err)
	}

	s := grpc.NewServer()

	/**注册接口服务
	 * 以定义proto时的service为单位注册，服务中可以有多个方法
	 * (proto编译时会为每个service生成Register***Server方法)
	 * 包.注册服务方法(gRpc服务实例，包含接口方法的结构体[指针])
	 */
	protoMsg.RegisterLoginServiceServer(s, &server{})

	reflection.Register(s)

	err = s.Serve(listen)
	if err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	con.Connect("root:youpassword@tcp(192.168.3.210:3306)/game?charset=utf8mb4&parseTime=True&loc=Local")
	initRpc()
}
