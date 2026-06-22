package cache

import (
	"context"
	"time"
)

type Provider interface {
	Set(ctx context.Context, key string, value any, expiration time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Delete(ctx context.Context, key string) error
	Close() error
	Ping(ctx context.Context) error
}