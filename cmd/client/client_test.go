package client

import (
	"context"
	"github.com/dnk90/adlive/pb"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewAdliveClient(t *testing.T) {
	client := NewAdliveClient("localhost:10000")
	resp, err := client.GetUserProfile(context.Background(), &pb.EmptyMessage{})
	require.NoError(t, err)
	print(resp)
}
