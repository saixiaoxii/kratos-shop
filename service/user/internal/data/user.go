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
	"gorm.io/gorm"
)

// User define the user struct
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
	hashedPassword := encrypt(u.Password) // 假设存在加密函数

	entUser, err := r.data.edb.User.
		Create().
		SetMobile(u.Mobile).
		SetNickname(u.NickName).
		SetPassword(hashedPassword).
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

// Password encryption
func encrypt(psd string) string {
	options := &password.Options{SaltLen: 16, Iterations: 10000, KeyLen: 32, HashFunction: sha512.New}
	salt, encodedPwd := password.Encode(psd, options)
	return fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
}
