package application

import (
	"context"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func sendResourceLocation(ctx context.Context, id string) error {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		uri := md.Get("requestURI")
		if len(uri) == 1 {
			err := grpc.SendHeader(ctx, metadata.Pairs(
				"location", strings.TrimSuffix(uri[0], "/")+id,
				"x-http-code", "201"),
			)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
