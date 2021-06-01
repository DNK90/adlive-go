package client

import (
	"context"
	l "github.com/dnk90/adlive/helpers/log"
	"github.com/dnk90/adlive/jwt"
	"github.com/dnk90/adlive/pb"
	"google.golang.org/grpc"
)

var ll = l.New()

type AdliveClient struct {
	pb.AdliveClient
}

func NewAdliveClient(address string) *AdliveClient {
	jwtCreds := jwt.NewToken("abc")
	conn, err := grpc.DialContext(context.Background(), address,
		grpc.WithInsecure(),
		grpc.WithPerRPCCredentials(jwtCreds),
	)
	if err != nil {
		ll.Fatal("Failed to dial Adlive service", l.Error(err))
	}
	c := pb.NewAdliveClient(conn)
	return &AdliveClient{c}
}
