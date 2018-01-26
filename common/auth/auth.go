package auth

import (
	"context"
	"fmt"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	restful "github.com/emicklei/go-restful"
	"github.com/laidingqing/dabanshan/common/clients"
	"github.com/laidingqing/dabanshan/pb"
)

var (
	// SecretKey JWT secret key
	SecretKey = "welcome to dabanshan"
	//JwtMiddleware jwt middleware
	JwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})
)

//AuthError Auth Error struct
type AuthError struct {
	ErrorCode int
	Reason    string
}

//AccountCredentials session struct
type AccountCredentials struct {
	Id          string `json:"id"`
	Username    string `json:"name"`
	AccountType string `json:"accountType"`
	AccessToken string `json:"accessToken"`
}

//UnauthenticatedError ...
func UnauthenticatedError() error {
	return &AuthError{
		ErrorCode: 401,
		Reason:    "Invalid username or password",
	}
}

//InvalidAccessTokenError token invalid
func InvalidAccessTokenError() error {
	return &AuthError{
		ErrorCode: 401,
		Reason:    "Invalid access token",
	}
}

func (err *AuthError) Error() string {
	return fmt.Sprintf("[Error]:%v: %v", err.ErrorCode, err.Reason)
}

//CreateJWT create a jwt token
func CreateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims

	return token.SignedString([]byte(SecretKey))
}

//GetAuthenticateHeader get Account Info by authenticate header
func GetAuthenticateHeader(r *restful.Request) (AccountCredentials, error) {
	token, err := request.ParseFromRequest(r.Request, request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		})
	if err == nil {
		res, err := clients.GetAccountClient().GetByToken(context.Background(), &pb.GetByTokenRequest{Token: token.Raw})
		if err != nil {
			return AccountCredentials{}, err
		}
		return AccountCredentials{
			Id:       res.Account.Id,
			Username: res.Account.Username,
		}, nil
	}
	return AccountCredentials{}, nil
}

//ValidateTokenMiddleware validate token with jwt
func ValidateTokenMiddleware(r *restful.Request) error {
	token, err := request.ParseFromRequest(r.Request, request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		})

	if err == nil {
		if token.Valid {
			return nil
		}
	}
	return InvalidAccessTokenError()
}
