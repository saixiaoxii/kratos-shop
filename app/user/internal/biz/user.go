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
	GetUser(context.Context, int64) (*User, error)
	ListUser(context.Context) ([]*User, error)
	CheckPassword(ctx context.Context, password, encryptedPassword string) (bool, error)
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

func (uc *UserUsecase) Get(ctx context.Context, id int64) (*User, error) {
	return uc.repo.GetUser(ctx, id)
}

func (uc *UserUsecase) List(ctx context.Context) ([]*User, error) {
	return uc.repo.ListUser(ctx)
}

func (uc *UserUsecase) CheckPassword(ctx context.Context, password, encryptedPassword string) (bool, error) {
	return uc.repo.CheckPassword(ctx, password, encryptedPassword)
}
