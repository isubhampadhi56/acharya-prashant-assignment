package logger

import (
	"log"

	"go.uber.org/zap"
)

var appLogger *zap.SugaredLogger
var auditLogger *zap.SugaredLogger

func InitializeAppLogger() *zap.SugaredLogger {
	if appLogger != nil {
		return appLogger
	}
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"app.log", "stdout"}
	logg, err := config.Build()

	if err != nil {
		log.Fatal(err)
	}

	appLogger = logg.Sugar()
	return appLogger
}

func InitializeAuditLogger() *zap.SugaredLogger {
	if auditLogger != nil {
		return auditLogger
	}
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"audit.log"}
	logg, err := config.Build()

	if err != nil {
		log.Fatal(err)
	}

	auditLogger = logg.Sugar()
	return auditLogger
}
