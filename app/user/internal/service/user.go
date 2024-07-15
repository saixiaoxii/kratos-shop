package service

import (
	"context"
	v1 "user/api/user/v1"
	"user/app/user/internal/biz"

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
		Role:     int(req.Role),
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
	u.log.Info(" Updating user ", " mobile:", req.Mobile, " nickname:", req.NickName, "Gender:", req.Gender)

	user, err := u.uc.Update(ctx, &biz.User{
		Mobile:   req.Mobile,
		Password: req.Password,
		NickName: req.NickName,
		Gender:   req.Gender,
		Role:     int(req.Role),
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

func (u *UserService) GetUser(ctx context.Context, req *v1.GetUserRequest) (*v1.GetUserReply, error) {
	u.log.Info(" Getting user ", " id:", req.Id)

	user, err := u.uc.Get(ctx, req.Id)
	if err != nil {
		u.log.Error("Failed to get user", "error", err)
		return nil, err
	}

	userInfoRsp := v1.GetUserReply{
		Id:       user.ID,
		Mobile:   user.Mobile,
		Password: user.Password,
		NickName: user.NickName,
	}

	return &userInfoRsp, nil
}

func (u *UserService) ListUser(ctx context.Context, req *v1.ListUserRequest) (*v1.ListUserReply, error) {
	u.log.Info("Listing users")

	users, err := u.uc.List(ctx)
	if err != nil {
		u.log.Error("Failed to list users", "error", err)
		return nil, err
	}

	userInfoRsp := &v1.ListUserReply{}
	for _, user := range users {
		userDetail := &v1.UserDetail{
			Id:       user.ID,
			Mobile:   user.Mobile,
			Password: user.Password,
			NickName: user.NickName,
			Birthday: user.Birthday,
			Gender:   user.Gender,
		}
		userInfoRsp.Users = append(userInfoRsp.Users, userDetail)
	}

	return userInfoRsp, nil
}
