# ent 入门

## ent 安装

```shell
go install entgo.io/ent/cmd/ent@latest

# 查看是否安装成功
ent
```

## ent 初始化

```shell
# 默认在 ./ent/schema 下创建文件,注意首字母要大写
ent new User

ent generate ./ent/schema
```

生成下列内容

```text
ent
├── client.go
├── config.go
├── context.go
├── ent.go
├── example_test.go
├── migrate
│   ├── migrate.go
│   └── schema.go
├── predicate
│   └── predicate.go
├── schema
│   └── user.go
├── tx.go
├── user
│   ├── user.go
│   └── where.go
├── user.go
├── user_create.go
├── user_delete.go
├── user_query.go
└── user_update.go
```

我们只需要修改`./ent/schema`中对应的表结构即可

在这里则是通过 ent 创建的`user.go`
```text
├── schema
│   └── user.go
```

下面是一个示例配置

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
        // id 列，go 代码中对应操作字段为 ID , 唯一
        field.Int64("id").StorageKey("ID").Unique(),
        // mobile 列，go 代码中对应操作字段为 Mobile，唯一且不能为空，最大长度为11，表示手机号码，用户唯一标识
        field.String("mobile").StorageKey("Mobile").Unique().NotEmpty().MaxLen(11).Comment("手机号码，用户唯一标识"),
        // password 列，go 代码中对应操作字段为 Password，不能为空，最大长度为100，表示用户密码，保存时需要注意是否加密
        field.String("password").StorageKey("Password").NotEmpty().MaxLen(100).Comment("用户密码的保存需要注意是否加密"),
        // nickname 列，go 代码中对应操作字段为 NickName，最大长度为25，表示用户昵称
        field.String("nickname").StorageKey("NickName").MaxLen(25).Comment("用户昵称"),
        // birthday 列，go 代码中对应操作字段为 Birthday，可以为空，表示用户的出生日期
        field.Time("birthday").StorageKey("Birthday").Nillable().Optional().Comment("出生日期"),
        // gender 列，go 代码中对应操作字段为 Gender，默认值为 "male"，最大长度为16，表示用户性别，female:女,male:男
        field.String("gender").StorageKey("Gender").Default("male").MaxLen(16).Comment("female:女,male:男"),
        // role 列，go 代码中对应操作字段为 Role，默认值为 1，表示用户角色，1:普通用户,2:管理员
        field.Int("role").StorageKey("Role").Default(1).Comment("1:普通用户,2:管理员"),
        // created_at 列，go 代码中对应操作字段为 CreatedAt，可以为空，表示记录的创建时间
        field.Time("created_at").StorageKey("add_time").Optional(),
        // updated_at 列，go 代码中对应操作字段为 UpdatedAt，可以为空，表示记录的更新时间
        field.Time("updated_at").StorageKey("update_time").Optional(),
        // is_deleted 列，go 代码中对应操作字段为 IsDeletedAt，默认值为 false，表示记录是否被删除
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

修改完成后需要重新生成对应文件

```shell
ent generate ./ent/schema
```

后续使用 ent 提供的方法进行增删查改即可