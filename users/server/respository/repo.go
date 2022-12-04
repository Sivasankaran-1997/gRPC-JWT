package respository

import (
	"fmt"
	"gRPC_jwt/users/server/database"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type User struct {
	ID       string `bson:"id"`
	Name     string `bson:"name"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
	Phoneno  string `bson:"phoneno"`
	Address  string `bson:"address"`
}

type UserJWTsigneDetails struct {
	Email string
	jwt.RegisteredClaims
}

func (user *User) CreateVaildate() error {
	if strings.TrimSpace(user.ID) == "" {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("User ID is Required"),
		)
	}

	if strings.TrimSpace(user.Name) == "" {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("User Name is Required"),
		)
	}

	if strings.TrimSpace(user.Email) == "" {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("User Email is Required"),
		)
	}

	if strings.TrimSpace(user.Password) == "" {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("User Password is Required"),
		)
	}

	if strings.TrimSpace(user.Phoneno) == "" {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("User Phoneno is Required"),
		)
	}

	if strings.TrimSpace(user.Address) == "" {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("User Address is Required"),
		)
	}
	return nil
}

func (user *User) LoginVaildate() error {
	if strings.TrimSpace(user.Email) == "" {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("User Email is Required"),
		)
	}

	if strings.TrimSpace(user.Password) == "" {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("User Password is Required"),
		)
	}
	return nil
}

func (user *User) Create() (*User, error) {

	var emailcount int64
	database.Instance.Model(&User{}).Where("email = ?", user.Email).Count(&emailcount)

	if emailcount > 0 {
		return nil, status.Errorf(
			codes.AlreadyExists,
			fmt.Sprintf("Email ID is Exists"),
		)
	}

	var emailPhoneno int64
	database.Instance.Model(&User{}).Where("phoneno = ?", user.Phoneno).Count(&emailPhoneno)

	if emailPhoneno > 0 {
		return nil, status.Errorf(
			codes.AlreadyExists,
			fmt.Sprintf("Phone No is Exists"),
		)
	}

	value := database.Instance.Create(&user)

	if value.Error != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Error on create user"),
		)
	}

	return user, nil
}

func (user *User) FindUser() (*User, error) {

	var emailcount int64
	database.Instance.Model(&User{}).Where("email = ?", user.Email).Count(&emailcount)

	if emailcount == 0 {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Email Not Found"),
		)
	}

	database.Instance.Find(&user, "email = ?", user.Email)

	return user, nil
}
