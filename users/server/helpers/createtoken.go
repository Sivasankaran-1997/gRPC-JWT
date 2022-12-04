package helpers

import (
	"fmt"
	"gRPC_jwt/users/server/respository"
	"os"

	"github.com/golang-jwt/jwt/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func CreateToken(claims *respository.UserJWTsigneDetails) (string, error) {

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		return "", status.Errorf(
			codes.Internal,
			fmt.Sprintf("Token Not Created"),
		)
	}

	return token, nil
}
