package example

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/zhtfirst/go-packages/logger"
)

const CSTLayout = "2006-01-02 15:04:05"
const ProjectAccessLogFile = "./example/logs/test-access.log"

func Logging() {
	// 方案一：
	//logging.Setup(false, "test")
	//
	//zap.S().Info("failed to fetch URL:")
	////zap.S().Errorf("offer_records error: %v", "err")
	//zap.S().Info("failed to fetch URL:", zap.String("url", "http://example.com"), zap.Int("attempt", 3), zap.Duration("backoff", 1))

	// 方案二：
	// 初始化 access logger
	accessLogger, err := logger.NewJSONLogger(
		logger.WithDisableConsole(),
		logger.WithField("domain", fmt.Sprintf("%s[%s]", "go-gin-api", "dev")),
		logger.WithTimeLayout(CSTLayout),
		logger.WithFileP(ProjectAccessLogFile),
	)
	if err != nil {
		panic(err)
	}

	accessLogger.Fatal("http server startup err", zap.Error(err))
}
