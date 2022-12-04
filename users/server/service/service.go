package service

import (
	"fmt"
	"gRPC_jwt/users/server/helpers"
	"gRPC_jwt/users/server/respository"
	"gRPC_jwt/users/server/utils"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/rs/xid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	UsersService UsersServiceInterface = &usersService{}
)

type UsersServiceInterface interface {
	CreateUser(data *respository.User) (*respository.User, error)
	LoginUser(data *respository.User) (string, error)
}

type usersService struct{}

func (s *usersService) CreateUser(data *respository.User) (*respository.User, error) {

	if err := data.CreateVaildate(); err != nil {
		return nil, err
	}

	uid := xid.New()
	data.ID = uid.String()

	pass := utils.HashPasswordMD5(data.Password)
	data.Password = pass

	result, err := data.Create()

	if err != nil {
		return nil, err

	}
	result.Password = ""

	return result, nil
}

func (s *usersService) LoginUser(data *respository.User) (string, error) {

	if err := data.LoginVaildate(); err != nil {
		return "", err
	}

	pass := data.Password

	result, insertErr := data.FindUser()
	if insertErr != nil {
		return "", insertErr
	}

	hpass := result.Password

	passcheck := utils.CheckHash(hpass, pass)

	if !passcheck {
		return "", status.Errorf(
			codes.Internal,
			fmt.Sprintf("Invalid Password"),
		)
	}

	userSignedStruct := &respository.UserJWTsigneDetails{
		data.Email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(10 * time.Minute)),
			Subject:   data.Email,
		},
	}

	tokenResp, tokenErr := helpers.CreateToken(userSignedStruct)

	if tokenErr != nil {
		return "", tokenErr
	}

	return tokenResp, nil
}

func GetEmail(email string) (*respository.User, error) {
	if strings.TrimSpace(email) == "" {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Email Required"),
		)
	}

	data := &respository.User{Email: email}
	result, insertErr := data.FindUser()
	if insertErr != nil {
		return nil, insertErr
	}

	result.Password = ""
	return result, nil

}
