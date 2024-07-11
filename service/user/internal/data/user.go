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
