# Go-kratos 框架商城微服务实战四

## 接口测试，这里采用 testify 进行断言测试

- 新建`/test/`目录
- 新建 /test/user_grpc_test.go
```go
package main

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	v1 "user/api/user/v1" // 导入 user client

	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/hashicorp/consul/api"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
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

```
- 新建 /test/user_http_test.go
```go
package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

var jwtKey = []byte("testKey")

func setupHTTPClient() *http.Client {
	return &http.Client{
		Timeout: time.Second * 10,
	}
}

func generateJWT() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 1).Unix(),
		"iat": time.Now().Unix(),
	})
	return token.SignedString(jwtKey)
}

func TestCreateUserHTTP(t *testing.T) {
	hClient := setupHTTPClient()

	reqBody, err := json.Marshal(map[string]string{
		"mobile":   "1381121120",
		"password": "password",
		"nickName": "nickname",
		"role":     "0",
	})
	assert.NoError(t, err)

	token, err := generateJWT()
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "http://127.0.0.1:8002/v1/users", bytes.NewBuffer(reqBody))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token) // 使用生成的JWT

	resp, err := hClient.Do(req)
	assert.NoError(t, err)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Error closing response body:", err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode, string(body))
	log.Printf("[http] CreateUser %+v\n", resp)
}

func TestUpdateUserHTTP(t *testing.T) {
	hClient := setupHTTPClient()
	reqBody, err := json.Marshal(map[string]string{
		"mobile":   "138110110",
		"password": "newpassword",
		"nickName": "newnickname",
		"gender":   "female",
		"role":     "2",
	})
	assert.NoError(t, err)

	token, err := generateJWT()
	assert.NoError(t, err)

	req, err := http.NewRequest("PUT", "http://127.0.0.1:8002/v1/users/2", bytes.NewBuffer(reqBody))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token) // 使用生成的JWT

	resp, err := hClient.Do(req)
	assert.NoError(t, err)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Error closing response body:", err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode, string(body))
	log.Printf("[http] UpdateUser %+v\n", resp)
}
```
