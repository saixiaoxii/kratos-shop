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