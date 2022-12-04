package controller

import (
	"context"
	"fmt"
	pb "gRPC_jwt/users/proto"
	"gRPC_jwt/users/server/dto"
	"gRPC_jwt/users/server/middleware"
	"gRPC_jwt/users/server/respository"
	"gRPC_jwt/users/server/service"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type userController struct{}

func NewUserControllerServer() pb.UserServiceServer {
	return userController{}
}

func (userController) CreateUser(ctx context.Context, req *pb.ProtoCreateRequest) (*pb.ProtoCreateReponse, error) {

	result := dto.NewUserRequest(req.GetReq())

	data, err := service.UsersService.CreateUser(result)

	if err != nil {
		return nil, err
	}

	response := dto.NewUserResponse(data)

	return response, nil
}

func (userController) Login(ctx context.Context, req *pb.ProtoLoginRequest) (*pb.ProtoLoginResponse, error) {

	result := dto.NewLoginRequest(req.ProtoEmail, req.ProtoPassword)

	data, err := service.UsersService.LoginUser(result)

	if err != nil {
		return nil, err
	}

	response := dto.NewLoginReponse(data)
	return response, nil
}

func (userController) GetUser(ctx context.Context, req *pb.ProtoGetRequest) (*pb.ProtoGetResponse, error) {

	md, ok := metadata.FromIncomingContext(ctx)
	var values []string
	if ok {
		values = md.Get("authorization")
		if strings.TrimSpace(values[0]) == "" {
			return nil, status.Errorf(
				codes.Internal,
				fmt.Sprintf("Invalid Metadata"),
			)
		}
	}

	result, resterr := middleware.UsergetValidate(values[0])

	if resterr != nil {
		return nil, resterr
	}

	response := dto.NewGetResponse(result.(*respository.User))
	return response, nil
}
