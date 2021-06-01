package jwt

import (
	"context"
	"google.golang.org/grpc/credentials"
)

type jwt struct {
	token string
}

func NewToken(token string) credentials.PerRPCCredentials {
	return &jwt{token: token}
}

func (j jwt) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": j.token,
	}, nil
}

func (j jwt) RequireTransportSecurity() bool {
	return false
}
