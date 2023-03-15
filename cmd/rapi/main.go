package main

import (
	"context"
	"fmt"
	"github.com/topklabs/rapi/internal/logging"
	"github.com/topklabs/rapi/pkg/proxy"
	"net/http"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := proxy.GetConfig()
	if err != nil {
		panic(err)
	}

	logger := logging.DefaultLogger()
	ctx = logging.NewContext(ctx, logger)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), proxy.NewProxy(ctx, cfg)); err != nil && err != context.Canceled {
		logger.Fatal(err)
	}
}
