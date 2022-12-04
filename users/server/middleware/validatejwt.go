package middleware

import (
	"fmt"
	"gRPC_jwt/users/server/service"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var User interface{}

func ValidateToken(tokens string) (interface{}, error) {
	token, _ := jwt.Parse(tokens, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		User = claims["sub"]
	} else if float64(time.Now().Unix()) > claims["exp"].(float64) && token.Valid {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("JWT Expired"),
		)
	} else {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Token Not Valid"),
		)
	}

	return User, nil
}

func UsergetValidate(tokens string) (interface{}, error) {

	tokenValue, err := ValidateToken(tokens)

	if err != nil {
		return nil, err
	}

	result, resterr := service.GetEmail(tokenValue.(string))

	if resterr != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("User Not Found"),
		)
	}

	return result, nil
}
