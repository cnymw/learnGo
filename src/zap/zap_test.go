package zap

import (
	"fmt"
	"go.uber.org/zap"
	"hash/crc32"
	"testing"
	"time"
)

func TestLog(t *testing.T) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger.Info("failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.String("url", "http://www.baidu.com"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)

	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", "http://www.baidu.com",
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", "http://www.baidu.com")

}

func Test1(t *testing.T) {
	fileid := "ecd5badd-cd51-417a-847d-f41fdbb30dcb11"
	appid := "822c6ba7-f7c6-4a56-0000-000000000001"
	uniqueid := int64(crc32.ChecksumIEEE([]byte(appid + fileid)))
	fmt.Println(uniqueid)
	fmt.Println(uniqueid % 1000)
}
