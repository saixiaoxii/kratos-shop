package main

import (
	"context"
	"log"
	"time"

	v1 "user/api/user/v1" // 导入 user client

	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/hashicorp/consul/api"
)

func main() {
    c := api.DefaultConfig()
    c.Address = "192.168.29.130:8500"
    c.Scheme = "http"
    consulCli, err := api.NewClient(c)
    if err != nil {
        panic(err)
    }
    r := consul.New(consulCli)

    // new grpc client
    conn, err := grpc.DialInsecure(
        context.Background(),
        grpc.WithEndpoint("discovery:///user"),
        grpc.WithDiscovery(r),
    )
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()
    gClient := v1.NewUserClient(conn) // 使用 UserClient

    // new http client
    hConn, err := http.NewClient(
        context.Background(),
        http.WithMiddleware(
            recovery.Recovery(),
        ),
        http.WithEndpoint("discovery:///user"),
        http.WithDiscovery(r),
    )
    if err != nil {
        log.Fatal(err)
    }
    defer hConn.Close()
    hClient := v1.NewUserHTTPClient(hConn) // 使用 UserHTTPClient

    for {
        time.Sleep(time.Second)
        callGRPC(gClient)
        callHTTP(hClient) // 调用 HTTP 客户端
    }
}

func callGRPC(client v1.UserClient) {
    // 这里调用 UserClient 的方法，例如 UpdateUser
    rsp, err := client.UpdateUser(context.Background(), &v1.UpdateUserRequest{
        Mobile:   "138110110",
        Password: "password",
        NickName: "nic11kna44me",
    })
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("[grpc] UpdateUser %+v\n", rsp)
}

func callHTTP(client v1.UserHTTPClient) {
    // 这里调用 UserHTTPClient 的方法，例如 UpdateUser
    rsp, err := client.UpdateUser(context.Background(), &v1.UpdateUserRequest{
        Mobile:   "138110110",
        Password: "password",
        NickName: "nic11kna44me",
    })
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("[http] UpdateUser %+v\n", rsp)
}