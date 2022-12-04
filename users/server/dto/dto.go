package dto

import (
	pb "gRPC_jwt/users/proto"
	"gRPC_jwt/users/server/respository"
)

func NewUserRequest(data *pb.ProtoUser) *respository.User {
	return &respository.User{
		ID:       data.ProtoID,
		Name:     data.ProtoName,
		Email:    data.ProtoEmail,
		Password: data.ProtoPassword,
		Phoneno:  data.ProtoPhoneno,
		Address:  data.ProtoAddress,
	}
}

func NewUserResponse(data *respository.User) *pb.ProtoCreateReponse {
	user := &pb.ProtoUser{
		ProtoID:       data.ID,
		ProtoName:     data.Name,
		ProtoEmail:    data.Email,
		ProtoPassword: data.Password,
		ProtoPhoneno:  data.Phoneno,
		ProtoAddress:  data.Address,
	}
	return &pb.ProtoCreateReponse{
		Res: user,
	}
}

func NewGetResponse(data *respository.User) *pb.ProtoGetResponse {
	user := &pb.ProtoUser{
		ProtoID:       data.ID,
		ProtoName:     data.Name,
		ProtoEmail:    data.Email,
		ProtoPassword: data.Password,
		ProtoPhoneno:  data.Phoneno,
		ProtoAddress:  data.Address,
	}
	return &pb.ProtoGetResponse{
		Res: user,
	}
}

func NewLoginRequest(userEmail string, userPassword string) *respository.User {
	return &respository.User{
		Email:    userEmail,
		Password: userPassword,
	}
}

func NewLoginReponse(token string) *pb.ProtoLoginResponse {
	return &pb.ProtoLoginResponse{Token: token}
}
