# Go-kratos 框架商城微服务实战一

## 准备工作，初始化项目目录

```bash
mkdir  -p kratos-shop/service
cd kratos-shop/service

kratos new user
cd user

kratos proto add api/user/v1/user.proto

kratos proto client api/user/v1/user.proto

kratos proto server api/user/v1/user.proto -t internal/service

go generate ./...
```

## 接口定义，在 `user.proto` 定义

```proto
syntax = "proto3";

package api.user.v1;

option go_package = "user/api/user/v1;v1";
option java_multiple_files = true;
option java_package = "api.user.v1";

service User {
	rpc CreateUser (CreateUserRequest) returns (CreateUserReply);
	rpc UpdateUser (UpdateUserRequest) returns (UpdateUserReply);
	rpc DeleteUser (DeleteUserRequest) returns (DeleteUserReply);
	rpc GetUser (GetUserRequest) returns (GetUserReply);
	rpc ListUser (ListUserRequest) returns (ListUserReply);
}

message CreateUserRequest {
	string nickName = 1;
	string password = 2;
	string mobile = 3;
}
message CreateUserReply {
	int64 id = 1;
	string password = 2;
	string mobile = 3;
	string nickName = 4;
	int64 birthday = 5;
	string gender = 6;
	int32 role = 7;
}

message UpdateUserRequest {}
message UpdateUserReply {}

message DeleteUserRequest {}
message DeleteUserReply {}

message GetUserRequest {}
message GetUserReply {}

message ListUserRequest {}
message ListUserReply {}
```

## 生成信息

```bash
make api

# 或者使用
kratos proto client api/user/v1/user.proto
```

## 修改 `ser/configs/config.yaml`, 注意调整配置内容为自己的配置

```yaml
server:
  http:
    addr: 0.0.0.0:8000
    timeout: 1s
  grpc:
    addr: 0.0.0.0:50051
    timeout: 1s
data:
  database:
    driver: mysql
    #这里的数据库配置修改为自己的数据库配置
    source: root:123456@tcp(192.168.29.130:3306)/shop_user?charset=utf8mb4&parseTime=True&loc=Local
  redis:
    addr: 192.168.29.130:6379
    dial_timeout: 1s
    read_timeout: 0.2s
    write_timeout: 0.2s
trace:
  endpoint: http://192.168.29.130:14268/api/traces
```

## 新建 `user/configs/registry.yaml`, 引入consul服务

```bash
# 这里引入了 consul 的服务注册与发现，先把配置加入进去
consul:
  address: 192.168.29.130:8500
  scheme: http
```

## 修改 `user/internal/conf/conf.proto` 配置文件

```go
syntax = "proto3";
package kratos.api;

option go_package = "user/internal/conf;conf";

import "google/protobuf/duration.proto";

message Bootstrap {
  Server server = 1;
  Data data = 2;
}

message Server {
  message HTTP {
    string network = 1;
    string addr = 2;
    google.protobuf.Duration timeout = 3;
  }
  message GRPC {
    string network = 1;
    string addr = 2;
    google.protobuf.Duration timeout = 3;
  }
  HTTP http = 1;
  GRPC grpc = 2;
}

message Data {
  message Database {
    string driver = 1;
    string source = 2;
  }
  message Redis {
    string network = 1;
    string addr = 2;
    string password = 3;
    int32 db = 4;
    google.protobuf.Duration dial_timeout = 5;
    google.protobuf.Duration read_timeout = 6;
    google.protobuf.Duration write_timeout = 7;
  }
  Database database = 1;
  Redis redis = 2;
}

message Registry {
  message Consul {
    string address = 1;
    string scheme = 2;
  }
  Consul consul = 1;
}

message Trace {
  string endpoint = 1;
}
```

## 生成对应 go 代码

```bash
make config
```

## docker 启动 consul 服务

```bash
docker run -d -p 8500:8500 -p 8300:8300 -p 8301:8301 -p 8302:8302 -p 8600:8600/udp consul consul agent -dev -client=0.0.0.0 --restart=always
```

## 服务代码

---

### mysql、redis 和新增用户的服务注册以及初始化，新增用户的功能代码在 
`data/user.go`

```go
package data

import (
	slog "log"
	"os"
	"time"
	"user/internal/conf"

	"github.com/go-redis/redis/extra/redisotel"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo, NewDB, NewRedis, NewUserRepo)

// Data .
type Data struct {
	db  *gorm.DB
	rdb *redis.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, db *gorm.DB, rdb *redis.Client) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{db: db, rdb: rdb}, cleanup, nil
}

// NewDB .
func NewDB(c *conf.Data) *gorm.DB {
	// 终端打印输入 sql 执行记录
	newLogger := logger.New(
		slog.New(os.Stdout, "\r\n", slog.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢查询 SQL 阈值
			Colorful:      true,        // 禁用彩色打印
			//IgnoreRecordNotFoundError: false,
			LogLevel: logger.Info, // Log lever
		},
	)

	db, err := gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{
		Logger:                                   newLogger,
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy:                           schema.NamingStrategy{
			//SingularTable: true, // 表名是否加 s
		},
	})

	if err != nil {
		log.Errorf("failed opening connection to sqlite: %v", err)
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&User{})
	if err != nil {
		panic(err)
	}

	return db
}

func NewRedis(c *conf.Data) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Redis.Addr,
		Password:     c.Redis.Password,
		DB:           int(c.Redis.Db),
		DialTimeout:  c.Redis.DialTimeout.AsDuration(),
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
	})
	rdb.AddHook(redisotel.TracingHook{})
	if err := rdb.Close(); err != nil {
		log.Error(err)
	}
	return rdb
}
```

### `biz/user.go` 定义 User 接口

```go
package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// User 定义返回数据结构体
type User struct {
	ID       int64
	Mobile   string
	Password string
	NickName string
	Birthday int64
	Gender   string
	Role     int
}

type UserRepo interface {
	CreateUser(context.Context, *User) (*User, error)
}

type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) Create(ctx context.Context, u *User) (*User, error) {
	return uc.repo.CreateUser(ctx, u)
}
```

### 在 `biz.go` 中注册服务

```go
package biz

import "github.com/google/wire"

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewGreeterUsecase, NewUserUsecase)
```

### 新增用户功能实现 `data/user.go`

```go
package data

import (
	"context"
	"crypto/sha512"
	"fmt"
	"time"
	"user/internal/biz"

	"github.com/anaskhan96/go-password-encoder"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

// 定义数据表结构体
type User struct {
    ID          int64      `gorm:"primarykey"`
    Mobile      string     `gorm:"index:idx_mobile;unique;type:varchar(11) comment '手机号码，用户唯一标识';not null"`
    Password    string     `gorm:"type:varchar(100);not null "` // 用户密码的保存需要注意是否加密
    NickName    string     `gorm:"type:varchar(25) comment '用户昵称'"`
    Birthday    *time.Time `gorm:"type:datetime comment '出生日期'"`
    Gender      string     `gorm:"column:gender;default:male;type:varchar(16) comment 'female:女,male:男'"`
    Role        int        `gorm:"column:role;default:1;type:int comment '1:普通用户,2:管理员'"`
    CreatedAt   time.Time  `gorm:"column:add_time"`
    UpdatedAt   time.Time  `gorm:"column:update_time"`
    DeletedAt   gorm.DeletedAt
    IsDeletedAt bool
}
type userRepo struct {
    data *Data
    log  *log.Helper
}

// NewUserRepo . 这里需要注意，上面 data 文件 wire 注入的是此方法，方法名不要写错了
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
    return &userRepo{
        data: data,
        log:  log.NewHelper(logger),
    }
}

// CreateUser .
func (r *userRepo) CreateUser(ctx context.Context, u *biz.User) (*biz.User, error) {
    var user User
    // 验证是否已经创建
    result := r.data.db.Where(&biz.User{Mobile: u.Mobile}).First(&user)
    if result.RowsAffected == 1 {
        return nil, status.Errorf(codes.AlreadyExists, "用户已存在")
    }

    user.Mobile = u.Mobile
    user.NickName = u.NickName
    user.Password = encrypt(u.Password) // 密码加密
    res := r.data.db.Create(&user)
    if res.Error != nil {
        return nil, status.Errorf(codes.Internal, res.Error.Error())
    }
    return &biz.User{
        ID:       user.ID,
        Mobile:   user.Mobile,
        Password: user.Password,
        NickName: user.NickName,
        Gender:   user.Gender,
        Role:     user.Role,
    }, nil
}

// Password encryption
func encrypt(psd string) string {
    options := &password.Options{SaltLen: 16, Iterations: 10000, KeyLen: 32, HashFunction: sha512.New}
    salt, encodedPwd := password.Encode(psd, options)
    return fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
}
```

### `service/user.go` 实现grpc接口请求

```go
package service

import (
	"context"
	v1 "user/api/user/v1"
	"user/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type UserService struct {
    v1.UnimplementedUserServer

    uc  *biz.UserUsecase
    log *log.Helper
}

// NewUserService new a greeter service.
func NewUserService(uc *biz.UserUsecase, logger log.Logger) *UserService {
    return &UserService{uc: uc, log: log.NewHelper(logger)}
}

// CreateUser create a user
func (u *UserService) CreateUser(ctx context.Context, req *v1.CreateUserRequest) (*v1.CreateUserReply, error) {
    user, err := u.uc.Create(ctx, &biz.User{
        Mobile:   req.Mobile,
        Password: req.Password,
        NickName: req.NickName,
    })
    if err != nil {
        return nil, err
    }

    userInfoRsp := v1.CreateUserReply{
        Id:       user.ID,
        Mobile:   user.Mobile,
        Password: user.Password,
        NickName: user.NickName,
        Gender:   user.Gender,
        Role:     int32(user.Role),
        Birthday: user.Birthday,
    }

    return &userInfoRsp, nil
}
```

### 在 `service.go` 中注册服务

```go
package service

import "github.com/google/wire"

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewGreeterService, NewUserService)
```

### `server/grpc.go` 中注册 grpc 服务

```go 
package server

import (
	v1 "user/api/helloworld/v1"
	v2 "user/api/user/v1"
	"user/internal/conf"
	"user/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, greeter *service.GreeterService, u *service.UserService,logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterGreeterServer(srv, greeter)
	v2.RegisterUserServer(srv, u)
	return srv
}
```


