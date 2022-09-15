package example

import (
	"github.com/zhtfirst/go-packages/logging"
	"go.uber.org/zap"
)

func Logging() {
	logging.Setup(false, "test")

	zap.S().Info("failed to fetch URL:")
	//zap.S().Errorf("offer_records error: %v", "err")
	zap.S().Info("failed to fetch URL:", zap.String("url", "http://example.com"), zap.Int("attempt", 3), zap.Duration("backoff", 1))
}
