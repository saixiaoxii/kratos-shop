package main

import (
	"context"
	v1 "kratos-shop/api/user/v1"
	"log"
	"os"
	"testing"
	"time"

	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"gopkg.in/yaml.v3"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/hashicorp/consul/api"
	"github.com/stretchr/testify/assert"
)

type Config struct {
	Consul struct {
		Address string `yaml:"address"`
		Scheme  string `yaml:"scheme"`
	} `yaml:"consul"`
}

func setupGRPCClient(t *testing.T) v1.UserClient {
	c := api.DefaultConfig()
	data, err := os.ReadFile("../configs/registry.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// 将读取的YAML数据解析到变量中
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	c.Address = config.Consul.Address
	c.Scheme = "http"
	consulCli, err := api.NewClient(c)
	assert.NoError(t, err)

	r := consul.New(consulCli)

	// new grpc client
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///user"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			metadata.Client(),
		),
	)

	assert.NoError(t, err)
	t.Cleanup(func() { conn.Close() })
	return v1.NewUserClient(conn) // 使用 UserClient
}

func TestCreateUser(t *testing.T) {
	gClient := setupGRPCClient(t)
	time.Sleep(time.Second)
	rsp, err := gClient.CreateUser(context.Background(), &v1.CreateUserRequest{
		Mobile:   "1381210110",
		Password: "password",
		NickName: "nickname",
	})
	assert.NoError(t, err)
	assert.NotNil(t, rsp)
	assert.Equal(t, "1381210110", rsp.Mobile)
	log.Printf("[grpc] CreateUser %+v\n", rsp)
}

func TestUpdateUser(t *testing.T) {
	gClient := setupGRPCClient(t)
	time.Sleep(time.Second)
	rsp, err := gClient.UpdateUser(context.Background(), &v1.UpdateUserRequest{
		Id:       1,
		Mobile:   "138110110",
		Password: "newpassword",
		NickName: "newnickname",
	})
	assert.NoError(t, err)
	assert.NotNil(t, rsp)
	assert.Equal(t, "138110110", rsp.Mobile)
	log.Printf("[grpc] UpdateUser %+v\n", rsp)
}
