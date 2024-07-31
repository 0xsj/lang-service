package user

import (
	"context"

	pb "github.com/0xsj/kakao/auth/proto"
)

type UserService struct {
	pb.UnimplementedAuthServiceServer
	
}

func Register(ctx context.Context, req *pb.AuthToken){
	return 
}