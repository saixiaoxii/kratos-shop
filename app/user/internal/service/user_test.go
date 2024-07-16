package service

import (
	"context"
	"testing"

	v1 "kratos-shop/api/user/v1"
	"kratos-shop/app/user/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUserRepo 是 biz.UserRepo 的模拟实现
type MockUserRepo struct {
	mock.Mock
}

func (m *MockUserRepo) CreateUser(ctx context.Context, user *biz.User) (*biz.User, error) {
	args := m.Called(ctx, user)
	return args.Get(0).(*biz.User), args.Error(1)
}

func (m *MockUserRepo) UpdateUser(ctx context.Context, user *biz.User) (*biz.User, error) {
	args := m.Called(ctx, user)
	return args.Get(0).(*biz.User), args.Error(1)
}

func (m *MockUserRepo) GetUser(ctx context.Context, id int64) (*biz.User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*biz.User), args.Error(1)
}

func (m *MockUserRepo) ListUser(ctx context.Context) ([]*biz.User, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*biz.User), args.Error(1)
}

func (m *MockUserRepo) CheckPassword(ctx context.Context, password, encryptedPassword string) (bool, error) {
	args := m.Called(ctx, password, encryptedPassword)
	return args.Bool(0), args.Error(1)
}

// 确保 MockUserRepo 实现了 biz.UserRepo 接口
var _ biz.UserRepo = (*MockUserRepo)(nil)

// 测试 UserService 的 ListUser 方法
func TestUserService_ListUser(t *testing.T) {
	mockRepo := new(MockUserRepo)
	logger := log.DefaultLogger
	userUsecase := biz.NewUserUsecase(mockRepo, logger)
	userService := NewUserService(userUsecase, logger)

	// 设置 UserRepo 的 ListUser 方法的预期行为和返回值
	expectedUsers := []*biz.User{
		{
			ID:       1,
			Mobile:   "12345678901",
			Password: "password",
			NickName: "nickname",
		},
	}
	mockRepo.On("ListUser", mock.Anything).Return(expectedUsers, nil)

	// 调用 UserService 的 ListUser 方法
	resp, err := userService.ListUser(context.Background(), &v1.ListUserRequest{})

	// 使用 testify 的 assert 函数进行断言
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Len(t, resp.Users, 1)
	assert.Equal(t, int64(1), resp.Users[0].Id)
	assert.Equal(t, "12345678901", resp.Users[0].Mobile)
	assert.Equal(t, "password", resp.Users[0].Password)
	assert.Equal(t, "nickname", resp.Users[0].NickName)

	mockRepo.AssertExpectations(t)
}

// 测试 UserService 的 CreateUser 方法
func TestUserService_CreateUser(t *testing.T) {
	mockRepo := new(MockUserRepo)
	logger := log.DefaultLogger
	userUsecase := biz.NewUserUsecase(mockRepo, logger)
	userService := NewUserService(userUsecase, logger)

	// 设置 UserRepo 的 CreateUser 方法的预期行为和返回值
	expectedUser := &biz.User{
		ID:       1,
		Mobile:   "12345678901",
		Password: "password",
		NickName: "nickname",
	}
	mockRepo.On("CreateUser", mock.Anything, mock.AnythingOfType("*biz.User")).Return(expectedUser, nil)

	// 调用 UserService 的 CreateUser 方法
	req := &v1.CreateUserRequest{
		Mobile:   "12345678901",
		Password: "password",
		NickName: "nickname",
		Role:     1,
	}
	resp, err := userService.CreateUser(context.Background(), req)

	// 使用 testify 的 assert 函数进行断言
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, int64(1), resp.Id)
	assert.Equal(t, "12345678901", resp.Mobile)
	assert.Equal(t, "password", resp.Password)
	assert.Equal(t, "nickname", resp.NickName)

	mockRepo.AssertExpectations(t)
}

// 测试 UserService 的 UpdateUser 方法
func TestUserService_UpdateUser(t *testing.T) {
	mockRepo := new(MockUserRepo)
	logger := log.DefaultLogger
	userUsecase := biz.NewUserUsecase(mockRepo, logger)
	userService := NewUserService(userUsecase, logger)

	// 设置 UserRepo 的 UpdateUser 方法的预期行为和返回值
	expectedUser := &biz.User{
		ID:       1,
		Mobile:   "12345678901",
		Password: "newpassword",
		NickName: "newnickname",
		Gender:   "male",
	}
	mockRepo.On("UpdateUser", mock.Anything, mock.AnythingOfType("*biz.User")).Return(expectedUser, nil)

	// 调用 UserService 的 UpdateUser 方法
	req := &v1.UpdateUserRequest{
		Mobile:   "12345678901",
		Password: "newpassword",
		NickName: "newnickname",
		Gender:   "male",
		Role:     1,
	}
	resp, err := userService.UpdateUser(context.Background(), req)

	// 使用 testify 的 assert 函数进行断言
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, int64(1), resp.Id)
	assert.Equal(t, "12345678901", resp.Mobile)
	assert.Equal(t, "newpassword", resp.Password)
	assert.Equal(t, "newnickname", resp.NickName)
	assert.Equal(t, "male", resp.Gender)

	mockRepo.AssertExpectations(t)
}

// 测试 UserService 的 GetUser 方法
func TestUserService_GetUser(t *testing.T) {
	mockRepo := new(MockUserRepo)
	logger := log.DefaultLogger
	userUsecase := biz.NewUserUsecase(mockRepo, logger)
	userService := NewUserService(userUsecase, logger)

	// 设置 UserRepo 的 GetUser 方法的预期行为和返回值
	expectedUser := &biz.User{
		ID:       1,
		Mobile:   "12345678901",
		Password: "password",
		NickName: "nickname",
	}
	mockRepo.On("GetUser", mock.Anything, int64(1)).Return(expectedUser, nil)

	// 调用 UserService 的 GetUser 方法
	resp, err := userService.GetUser(context.Background(), &v1.GetUserRequest{Id: 1})

	// 使用 testify 的 assert 函数进行断言
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, int64(1), resp.Id)
	assert.Equal(t, "12345678901", resp.Mobile)
	assert.Equal(t, "password", resp.Password)
	assert.Equal(t, "nickname", resp.NickName)

	mockRepo.AssertExpectations(t)
}

// 测试 UserService 的 CheckPassword 方法
func TestUserService_CheckPassword(t *testing.T) {
	mockRepo := new(MockUserRepo)
	logger := log.DefaultLogger
	userUsecase := biz.NewUserUsecase(mockRepo, logger)
	userService := NewUserService(userUsecase, logger)

	// 设置 UserRepo 的 CheckPassword 方法的预期行为和返回值
	mockRepo.On("CheckPassword", mock.Anything, "password", "encryptedPassword").Return(true, nil)

	// 调用 UserService 的 CheckPassword 方法
	resp, err := userService.CheckPassword(context.Background(), &v1.PasswordCheckInfo{
		Password:          "password",
		EncryptedPassword: "encryptedPassword",
	})

	// 使用 testify 的 assert 函数进行断言
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.True(t, resp.Success)

	mockRepo.AssertExpectations(t)
}
