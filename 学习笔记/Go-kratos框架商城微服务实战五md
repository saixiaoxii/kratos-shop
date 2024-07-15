# Go-kratos 框架商城微服务实战五

## 新增一个checkpassword的方法
这里就不提供修改部分，可以直接去[github]()仓库看源码
- 修改接口
- biz层新增方法
- data层新增方法
- service层新增方法

## 单元测试
这里对`data`层进行单元测试
- 新建一个/data/user_test.go
```go
package data

import (
	"context"
	"os"
	"testing"

	"kratos-shop/app/user/internal/biz"
	"kratos-shop/app/user/internal/data/ent/enttest"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func setupTestData(t *testing.T) (*Data, func()) {
	// 在内存中模拟连接
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	// 也可以采用mysql进行连接, 无法在内存中模拟操作, 需要连接具体的数据库 更换的时候记得更换导入的驱动
	// client := enttest.Open(t, "mysql", "user:password@tcp(localhost:3306)/dbname?parseTime=True")
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	data := &Data{
		edb: client,
		rdb: rdb,
	}
	cleanup := func() {
		client.Close()
		rdb.Close()
	}
	return data, cleanup
}

func TestCreateUser(t *testing.T) {
	data, cleanup := setupTestData(t)
	defer cleanup()

	logger := log.NewStdLogger(os.Stdout)
	repo := NewUserRepo(data, logger)

	ctx := context.Background()
	user := &biz.User{
		Mobile:   "13811211501",
		Password: "password",
		NickName: "nickname",
		Role:     1,
	}

	createdUser, err := repo.CreateUser(ctx, user)
	assert.NoError(t, err)
	assert.NotNil(t, createdUser)
	assert.Equal(t, user.Mobile, createdUser.Mobile)
}

func TestUpdateUser(t *testing.T) {
	data, cleanup := setupTestData(t)
	defer cleanup()

	logger := log.NewStdLogger(os.Stdout)
	repo := NewUserRepo(data, logger)

	ctx := context.Background()
	user := &biz.User{
		Mobile:   "13811211501",
		Password: "password",
		NickName: "nickname",
		Role:     1,
	}

	createdUser, err := repo.CreateUser(ctx, user)
	assert.NoError(t, err)
	assert.NotNil(t, createdUser)

	updatedUser := &biz.User{
		Mobile:   "13811211501",
		Password: "newpassword",
		NickName: "newnickname",
		Role:     2,
	}

	updatedUser, err = repo.UpdateUser(ctx, updatedUser)
	assert.NoError(t, err)
	assert.NotNil(t, updatedUser)
	assert.Equal(t, "newnickname", updatedUser.NickName)
}

func TestGetUser(t *testing.T) {
	data, cleanup := setupTestData(t)
	defer cleanup()

	logger := log.NewStdLogger(os.Stdout)
	repo := NewUserRepo(data, logger)

	ctx := context.Background()
	user := &biz.User{
		Mobile:   "13811211501",
		Password: "password",
		NickName: "nickname",
		Role:     1,
	}

	createdUser, err := repo.CreateUser(ctx, user)
	assert.NoError(t, err)
	assert.NotNil(t, createdUser)

	fetchedUser, err := repo.GetUser(ctx, createdUser.ID)
	assert.NoError(t, err)
	assert.NotNil(t, fetchedUser)
	assert.Equal(t, createdUser.ID, fetchedUser.ID)
}

func TestListUser(t *testing.T) {
	data, cleanup := setupTestData(t)
	defer cleanup()

	logger := log.NewStdLogger(os.Stdout)
	repo := NewUserRepo(data, logger)

	ctx := context.Background()
	user1 := &biz.User{
		Mobile:   "13811211501",
		Password: "password",
		NickName: "nickname1",
		Role:     1,
	}
	user2 := &biz.User{
		Mobile:   "13811211502",
		Password: "password",
		NickName: "nickname2",
		Role:     2,
	}

	_, err := repo.CreateUser(ctx, user1)
	assert.NoError(t, err)
	_, err = repo.CreateUser(ctx, user2)
	assert.NoError(t, err)

	users, err := repo.ListUser(ctx)
	assert.NoError(t, err)
	assert.Len(t, users, 2)
}

func TestCheckPassword(t *testing.T) {
	data, cleanup := setupTestData(t)
	defer cleanup()

	logger := log.NewStdLogger(os.Stdout)
	repo := NewUserRepo(data, logger)

	ctx := context.Background()
	user := &biz.User{
		Mobile:   "13811211501",
		Password: "password",
		NickName: "nickname",
		Role:     1,
	}

	createdUser, err := repo.CreateUser(ctx, user)
	assert.NoError(t, err)
	assert.NotNil(t, createdUser)

	valid, err := repo.CheckPassword(ctx, "password", createdUser.Password)
	assert.NoError(t, err)
	assert.True(t, valid)

	invalid, err := repo.CheckPassword(ctx, "wrongpassword", createdUser.Password)
	assert.NoError(t, err)
	assert.False(t, invalid)
}

```
可以对data层进行测试

后续补全biz层和service层