package main

import (
	"fmt"
	"log"

	"go.uber.org/zap/zapcore"

	"github.com/open-feature/flagd/pkg/logger"
)

func main() {
	l, err := logger.NewZapLogger(zapcore.DebugLevel, "console")
	if err != nil {
		log.Fatalf("initialize zap logger: %v", err)
	}

	fmt.Println("foo")

	l.Info("kube-flagd-proxy")
}
