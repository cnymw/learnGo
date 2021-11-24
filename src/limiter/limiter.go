package limiter

import (
	"context"
	"errors"
	"fmt"
	libredis "github.com/go-redis/redis/v8"
	"github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/store/memory"
	"github.com/ulule/limiter/v3/drivers/store/redis"
	"time"
)

type Adapter struct {
	cli libredis.UniversalClient
}

func (a *Adapter) Get(ctx context.Context, key string) *libredis.StringCmd {
	fmt.Printf("redis Get key: %s\n", key)
	return a.cli.Get(ctx, key)
}

func (a *Adapter) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *libredis.StatusCmd {
	fmt.Printf("redis Set key : %s\n", key)
	return a.cli.Set(ctx, key, value, expiration)
}

func (a *Adapter) Watch(ctx context.Context, handler func(*libredis.Tx) error, keys ...string) error {
	fmt.Printf("redis Watch key : %v\n", keys)
	withPrefixKeys := make([]string, 0, len(keys))
	for _, k := range keys {
		withPrefixKeys = append(withPrefixKeys, k)
	}
	return a.cli.Watch(ctx, handler, withPrefixKeys...)
}

func (a *Adapter) Del(ctx context.Context, keys ...string) *libredis.IntCmd {
	fmt.Printf("redis Watch key : %v\n", keys)
	withPrefixKeys := make([]string, 0, len(keys))
	for _, k := range keys {
		withPrefixKeys = append(withPrefixKeys, k)
	}
	return a.cli.Del(ctx, withPrefixKeys...)
}

func (a *Adapter) SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) *libredis.BoolCmd {
	fmt.Printf("redis SetNX key : %s\n", key)
	return a.cli.SetNX(ctx, key, value, expiration)
}

func (a *Adapter) EvalSha(ctx context.Context, script string, keys []string, args ...interface{}) *libredis.Cmd {
	fmt.Printf("redis EvalSha keys : %v\n", keys)
	withPrefixKeys := make([]string, 0, len(keys))
	for _, k := range keys {
		withPrefixKeys = append(withPrefixKeys, k)
	}
	return a.cli.EvalSha(ctx, script, withPrefixKeys, args...)
}

func (a *Adapter) ScriptLoad(ctx context.Context, script string) *libredis.StringCmd {
	fmt.Printf("redis ScriptLoad script : %s\n", script)
	return a.cli.ScriptLoad(ctx, script)
}

func NewAdapter(r libredis.UniversalClient) *Adapter {
	return &Adapter{cli: r}
}

type Config struct {
	TimeOut    time.Duration
	SkipDemote bool //发生错误，直接跳过处理，不降级处理
	Key        func() (string, error)
	Period     int64
	Limit      int64
}

type Endpoint func(ctx context.Context) error

type Middleware func() Endpoint

func NewMiddleware(rds libredis.UniversalClient, cfg Config) Middleware {
	if rds == nil {
		panic("throttled: redis is nil")
	}

	store, err := redis.NewStore(NewAdapter(rds))
	if err != nil {
		panic(err)
	}

	rate := limiter.Rate{
		Period: time.Duration(cfg.Period * int64(time.Second)),
		Limit:  cfg.Limit,
	}

	redisLimiter := limiter.New(store, rate)

	var mLimiter *limiter.Limiter
	if cfg.SkipDemote == false {
		//降级处理
		mStore := memory.NewStore()
		mLimiter = limiter.New(mStore, rate)
	}

	return func() Endpoint {
		return func(ctx context.Context) error {
			key, err := cfg.Key()
			if err != nil {
				return err
			}
			if len(key) == 0 {
				return nil
			}

			timeOutCtx, cancel := context.WithTimeout(ctx, cfg.TimeOut)
			defer cancel()

			limitCtx, err := redisLimiter.Get(timeOutCtx, key)
			if err != nil {
				fmt.Printf("redisLimiter.Get err:%v\n", err)
				if cfg.SkipDemote {
					limitCtx = limiter.Context{Reached: false}
				} else {
					limitCtx, err = mLimiter.Get(ctx, key)
					if err != nil {
						//发生其他错误，记录，不阻塞接口
						return err
					}
				}
			}
			fmt.Printf("context %v\n", limitCtx)
			if !limitCtx.Reached {
				return nil
			} else {
				return errors.New("the caller has been throttled")
			}
		}
	}
}
