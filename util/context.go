package util

import (
	"context"

	"github.com/quincy0/qpro/qTrace"
)

func InitContext() context.Context {
	return qTrace.InitContextWithTrace(context.Background(), RandStr(32), RandStr(16))
}

func InitContextWithSameTrace(ctx context.Context) context.Context {
	return qTrace.InitContextWithTrace(context.Background(), qTrace.TraceIdFromContext(ctx), RandStr(16))
}
