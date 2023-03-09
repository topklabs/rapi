package main

import (
	"context"
	"fmt"
	"github.com/topklabs/rapi/pkg/proxy"
	"log"
	"net/http"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := proxy.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), proxy.NewProxy(ctx, cfg)); err != nil && err != context.Canceled {
		log.Fatal(err)
	}
}
