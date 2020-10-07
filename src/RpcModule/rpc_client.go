/**
* @Author: Jason
* @Date: 2020/9/29 17:38
* @File : rpc_client
* @Software: GoLand
**/

package RpcModule

import (
	protoMsg "dygame/message"
	"log"

	"google.golang.org/grpc"
)

var DbRpcClient protoMsg.LoginServiceClient

func InitRpc() {
	conn, err := grpc.Dial("127.0.0.1:18028", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return
	}

	DbRpcClient = protoMsg.NewLoginServiceClient(conn)
}
