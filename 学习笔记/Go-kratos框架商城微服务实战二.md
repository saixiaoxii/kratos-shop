# Go-kratos 框架商城微服务实战二

## `gorm` 转成 `ent`

综合考虑之后，采用`ent`作为数据库 ORM 映射工具

- ent命令建表
```shell
#默认在 ./ent/schema 下生成文件
ent new User 
```
- 设置表结构
```go
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").StorageKey("ID").Unique(),
		field.String("mobile").StorageKey("Mobile").Unique().NotEmpty().MaxLen(11).Comment("手机号码，用户唯一标识"),
		field.String("password").StorageKey("Password").NotEmpty().MaxLen(100).Comment("用户密码的保存需要注意是否加密"),
		field.String("nickname").StorageKey("NickName").MaxLen(25).Comment("用户昵称"),
		field.Time("birthday").StorageKey("Birthday").Nillable().Optional().Comment("出生日期"),
		field.String("gender").StorageKey("Gender").Default("male").MaxLen(16).Comment("female:女,male:男"),
		field.Int("role").StorageKey("Role").Default(1).Comment("1:普通用户,2:管理员"),
		field.Time("created_at").StorageKey("add_time").Optional(),
		field.Time("updated_at").StorageKey("update_time").Optional(),
		field.Bool("is_deleted").StorageKey("IsDeletedAt").Default(false),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

// Indexes of the User. 
func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("mobile").Unique(),
	}
}
```
- 调整 /internal/data/data.go
```go
package data

import (
	"context"
	"user/internal/conf"
	"user/internal/data/ent"

	"github.com/go-redis/redis/extra/redisotel"
	"github.com/go-redis/redis/v8"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo,  NeweDB, NewRedis, NewUserRepo)

// Data .
type Data struct {
	edb *ent.Client
	rdb *redis.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, edb *ent.Client, rdb *redis.Client) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{edb: edb, rdb: rdb}, cleanup, nil
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

func NeweDB(c *conf.Data) *ent.Client {
    client, err := ent.Open(c.Database.Driver, c.Database.Source)
    if err != nil {
        log.Errorf("failed opening connection to %s: %v", c.Database.Driver, err)
        panic("failed to connect database")
    }

    // 执行数据库模式创建或迁移
    ctx := context.Background()
    if err := client.Schema.Create(ctx); err != nil {
        log.Fatalf("failed creating schema resources: %v", err)
    }

    return client
}
```

## 新增 Update

- 修改 /internal/data/user.go, 这里顺便新增UpdateUser,记得修改接口文件（user.proto）和/internal/biz/user.go
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
	UpdateUser(context.Context, *User) (*User, error)
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

func (uc *UserUsecase) Update(ctx context.Context, u *User) (*User, error) {
	return uc.repo.UpdateUser(ctx, u)
}
```

```go 
package data

import (
	"context"
	"crypto/sha512"
	"errors"
	"fmt"
	"time"
	"user/internal/biz"
	"user/internal/data/ent/user"

	"github.com/anaskhan96/go-password-encoder"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UserRepo .
type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo wired userRepo
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreateUser .
func (r *userRepo) CreateUser(ctx context.Context, u *biz.User) (*biz.User, error) {
	// 检查 r.data 和 r.data.edb 不为 nil
	if r.data == nil || r.data.edb == nil {
		return nil, errors.New("data 或 edb 未初始化")
	}

	// 检查 u 不为 nil
	if u == nil {
		return nil, errors.New("传入的用户对象为 nil")
	}
	// 检查用户是否已存在
	exists, err := r.data.edb.User.
		Query().
		Where(user.MobileEQ(u.Mobile)).
		Exist(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "查询用户失败: %v", err)
	}
	if exists {
		return nil, status.Errorf(codes.AlreadyExists, "用户已存在")
	}

	// 创建新用户
	hashedPassword := encrypt(u.Password) 

	entUser, err := r.data.edb.User.
		Create().
		SetMobile(u.Mobile).
		SetNickname(u.NickName).
		SetPassword(hashedPassword).SetRole(u.Role).
		Save(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "创建用户失败: %v", err)
	}

	// 构造并返回业务层用户对象
	return &biz.User{
		ID:       entUser.ID,
		Mobile:   entUser.Mobile,
		Password: entUser.Password,
		NickName: entUser.Nickname,
		Gender:   entUser.Gender,
		Role:     entUser.Role,
	}, nil
}

func (r *userRepo) UpdateUser(ctx context.Context, u *biz.User) (*biz.User, error) {
	// 检查 r.data 和 r.data.edb 不为 nil
	if r.data == nil || r.data.edb == nil {
		return nil, errors.New("data 或 edb 未初始化")
	}

	// 检查 u 不为 nil
	if u == nil {
		return nil, errors.New("传入的用户对象为 nil")
	}

	// 检查用户是否存在
	entUser, err := r.data.edb.User.
		Query().
		Where(user.MobileEQ(u.Mobile)).
		First(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "查询用户失败: %v", err)
	}
	if entUser == nil {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}

	// 更新用户信息
	update := r.data.edb.User.UpdateOneID(entUser.ID)
	if u.NickName != "" {
		update.SetNickname(u.NickName)
	}
	if u.Password != "" {
		hashedPassword := encrypt(u.Password)
		update.SetPassword(hashedPassword)
	}
	if u.Gender != "" {
		update.SetGender(u.Gender)
	}
	if u.Role != 0 {
		update.SetRole(u.Role)
	}
	update.SetUpdatedAt(time.Now())
	entUser, err = update.Save(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "更新用户失败: %v", err)
	}

	// 构造并返回业务层用户对象
	return &biz.User{
		ID:       entUser.ID,
		Mobile:   entUser.Mobile,
		Password: entUser.Password,
		NickName: entUser.Nickname,
		Gender:   entUser.Gender,
		Role:     entUser.Role,
	}, nil
}

// Password encryption
func encrypt(psd string) string {
	options := &password.Options{SaltLen: 16, Iterations: 10000, KeyLen: 32, HashFunction: sha512.New}
	salt, encodedPwd := password.Encode(psd, options)
	return fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
}
```
- 接口文件修改
```proto
syntax = "proto3";

package api.user.v1;

option go_package = "user/api/user/v1;v1";
option java_multiple_files = true;
option java_package = "api.user.v1";
import "google/api/annotations.proto";
service User {
	rpc CreateUser (CreateUserRequest) returns (CreateUserReply) {
        option (google.api.http) = {
            post: "/v1/users"
            body: "*"
        };
    }
    rpc UpdateUser (UpdateUserRequest) returns (UpdateUserReply) {
        option (google.api.http) = {
            put: "/v1/users/{id}"
            body: "*"
        };
    }
    rpc DeleteUser (DeleteUserRequest) returns (DeleteUserReply) {
        option (google.api.http) = {
            delete: "/v1/users/{id}"
        };
    }
    rpc GetUser (GetUserRequest) returns (GetUserReply) {
        option (google.api.http) = {
            get: "/v1/users/{id}"
        };
    }
    rpc ListUser (ListUserRequest) returns (ListUserReply) {
        option (google.api.http) = {
            get: "/v1/users"
        };
    }
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

message UpdateUserRequest {
	int64 id = 1;
	string password = 2;
	string mobile = 3;
	string nickName = 4;
	int64 birthday = 5;
	string gender = 6;
	int32 role = 7;
}
message UpdateUserReply {
    int64 id = 1;
	string password = 2;
	string mobile = 3;
	string nickName = 4;
	int64 birthday = 5;
	string gender = 6;
	int32 role = 7;
}

message DeleteUserRequest {
	int64 id = 1;
}
message DeleteUserReply {}

message GetUserRequest {
	int64 id = 1;
}
message GetUserReply {}

message ListUserRequest {}
message ListUserReply {}
```

- 修改对应 service 层, /internal/service/user.go
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
	// 记录请求信息，注意避免记录敏感信息如密码
	u.log.Info(" Creating user ", " mobile:", req.Mobile, " nickname:", req.NickName)

	user, err := u.uc.Create(ctx, &biz.User{
		Mobile:   req.Mobile,
		Password: req.Password,
		NickName: req.NickName,
	})
	if err != nil {
		u.log.Error("Failed to create user", "error", err)
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

func (u *UserService) UpdateUser(ctx context.Context, req *v1.UpdateUserRequest) (*v1.UpdateUserReply, error) {
	u.log.Info(" Updating user ", " mobile:", req.Mobile, " nickname:", req.NickName, "Gender:",req.Gender)
 
	user, err := u.uc.Update(ctx, &biz.User{
		Mobile:   req.Mobile,
		Password: req.Password,
		NickName: req.NickName,
		Gender:   req.Gender,
		Role: int(req.Role),
	})
	if err != nil {
		u.log.Error("Failed to update user", "error", err)
		return nil, err
	}

	userInfoRsp := v1.UpdateUserReply{
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

