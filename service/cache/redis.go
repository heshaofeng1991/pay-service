package cache

import (
	"context"
	"pay-service/global"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
)

// LockContext 互斥锁 控制并发
func LockContext(ctx context.Context, sync *redsync.Redsync, name string, t time.Duration, option ...redsync.Option) (mutex *redsync.Mutex, err error) {
	if len(option) == 0 {
		option = append(option, redsync.WithExpiry(t))
	}

	mux := sync.NewMutex(name, option...)

	return mux, mux.LockContext(ctx)
}

func RedisClient(ctx context.Context) *redis.Client {
	return global.Srv.Redis.Client.WithContext(ctx)
}
