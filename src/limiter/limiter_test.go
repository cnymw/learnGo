package limiter

import (
	"context"
	libredis "github.com/go-redis/redis/v8"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestNewMiddleware(t *testing.T) {
	Convey("limiter", t, func() {
		rds := libredis.NewUniversalClient(&libredis.UniversalOptions{Addrs: []string{"localhost:6379"}})
		cfg := Config{
			Period: 1,
			Limit:  10,
			Key: func() (string, error) {
				return "test", nil
			},
			TimeOut: 200 * time.Millisecond,
		}

		mw := NewMiddleware(rds, cfg)
		endpoint := mw()
		for i := 0; i < 10; i++ {
			err := endpoint(context.Background())
			So(err, ShouldEqual, nil)
		}
		err := endpoint(context.Background())
		if err != nil {
			So(err.Error(), ShouldEqual, "the caller has been throttled")
		}
	})
}
