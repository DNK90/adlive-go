package helpers

import (
	"context"
	l "github.com/dnk90/adlive/helpers/log"
	"google.golang.org/grpc/metadata"
	"strconv"
)

var ll = l.New()

func GetUserIdFromContext(ctx context.Context) int {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		ll.S.Warnw("cannot get metadata from context")
		return 0
	}
	if md.Get("id") == nil {
		ll.S.Warnw("cannot find id in metadata")
		return 0
	}
	id, err := strconv.Atoi(md.Get("id")[0])
	if err != nil {
		ll.S.Errorw("error while convert id from metadata to int", "err", err)
		return 0
	}
	return id
}
