package util

import (
	"context"

	"github.com/google/uuid"
)

type ContextKey string

func (k ContextKey) String() string {
	return string(k)
}

func GetTraceIDFromContext(ctx context.Context) string {
	return ctx.Value(ContextKey("trace_id")).(string)
}

func SetTraceIDWithContext(ctx context.Context) context.Context {
	traceId := uuid.New().String()
	return context.WithValue(ctx, ContextKey("trace_id"), traceId)
}
