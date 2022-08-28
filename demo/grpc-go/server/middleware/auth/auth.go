package auth

import (
	"context"
	"errors"
	"fmt"
	"log"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
)

type TokenInfo struct {
	ID    string
	Roles []string
}

func AuthInterceptor(ctx context.Context) (context.Context, error) {
	fmt.Println("test middle auth AuthInterceptor")
	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	pr, ok := peer.FromContext(ctx)

	fmt.Println("log pr: ", pr, "ok: ", ok)
	if err != nil {
		log.Println("get bearer: error")
		return nil, err
	}

	tokenInfo, err := parseToken(token)
	if err != nil {
	  return nil, grpc.Errorf(codes.Unauthenticated, " %v", err)
	}
	nexCtx := context.WithValue(ctx, tokenInfo.ID, tokenInfo)
	return nexCtx, nil
}

func parseToken(token string) (TokenInfo, error) {
	var tokenInfo TokenInfo
	if token == "grpc.auth.token" {
		tokenInfo.ID = "1"
		tokenInfo.Roles = []string{"admin"}
		return tokenInfo, nil
	}
	return tokenInfo, errors.New("Token invalid")
}

