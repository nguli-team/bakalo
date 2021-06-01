package helper

import (
	"context"
)

var IPContextKey = "request-ip"

func GetRequestIP(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	if ip, ok := ctx.Value(IPContextKey).(string); ok {
		return ip
	}
	return ""
}
