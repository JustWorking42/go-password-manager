package auth

import (
	"context"
	"errors"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Auth struct {
	secret string
}

func NewAuth(secret string) *Auth { return &Auth{secret: secret} }

func (a *Auth) NewToken(userId primitive.ObjectID) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = userId

	tokenString, err := token.SignedString([]byte(a.secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (a *Auth) ParseToken(tokenString string) (primitive.ObjectID, error) {
	token, err := parseToken(tokenString, a.secret)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return primitive.ObjectID{}, err
	}
	id, ok := claims["id"].(string)
	if !ok {
		return primitive.ObjectID{}, errors.New("invalid token")
	}

	return primitive.ObjectIDFromHex(id)
}

func (a *Auth) OnlyAuthorizedInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "metadata is not provided")

	}

	tokenString, ok := md["authorization"]
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "authorization is not provided")
	}

	id, err := a.ParseToken(tokenString[0])
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "authorization is not valid")
	}
	newCtx := metadata.AppendToOutgoingContext(ctx, "id", id.Hex())

	return handler(newCtx, req)
}

func parseToken(tokenString, secret string) (*jwt.Token, error) {
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}