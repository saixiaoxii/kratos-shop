package service

import (
	"context"

	pb "kratos-shop/api/shop/v1"

	"google.golang.org/protobuf/types/known/emptypb"
)

type ShopService struct {
	pb.UnimplementedShopServer
}

func NewShopService() *ShopService {
	return &ShopService{}
}

func (s *ShopService) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterReply, error) {
	return &pb.RegisterReply{}, nil
}
func (s *ShopService) Login(ctx context.Context, req *pb.LoginReq) (*pb.RegisterReply, error) {
	return &pb.RegisterReply{}, nil
}
func (s *ShopService) Captcha(ctx context.Context, req *emptypb.Empty) (*pb.CaptchaReply, error) {
	return &pb.CaptchaReply{}, nil
}
func (s *ShopService) Detail(ctx context.Context, req *emptypb.Empty) (*pb.UserDetailResponse, error) {
	return &pb.UserDetailResponse{}, nil
}
func (s *ShopService) CreateAddress(ctx context.Context, req *pb.CreateAddressReq) (*pb.AddressInfo, error) {
	return &pb.AddressInfo{}, nil
}
func (s *ShopService) AddressListByUid(ctx context.Context, req *emptypb.Empty) (*pb.ListAddressReply, error) {
	return &pb.ListAddressReply{}, nil
}
func (s *ShopService) UpdateAddress(ctx context.Context, req *pb.UpdateAddressReq) (*pb.CheckResponse, error) {
	return &pb.CheckResponse{}, nil
}
func (s *ShopService) DefaultAddress(ctx context.Context, req *pb.AddressReq) (*pb.CheckResponse, error) {
	return &pb.CheckResponse{}, nil
}
func (s *ShopService) DeleteAddress(ctx context.Context, req *pb.AddressReq) (*pb.CheckResponse, error) {
	return &pb.CheckResponse{}, nil
}
