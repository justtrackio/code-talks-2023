package main

import (
	"context"
	"time"

	"github.com/justtrackio/gosoline/pkg/cfg"
	"github.com/justtrackio/gosoline/pkg/kernel"
	"github.com/justtrackio/gosoline/pkg/log"
)

type HelloWorld struct {
	logger log.Logger
	ticker *time.Ticker
}

func NewHelloWorld(ctx context.Context, config cfg.Config, logger log.Logger) (kernel.Module, error) {
	tickerInterval := config.GetDuration("ticker_interval")

	module := &HelloWorld{
		logger: logger.WithChannel("hello-world"),
		ticker: time.NewTicker(tickerInterval),
	}

	return module, nil
}

func (h HelloWorld) Run(ctx context.Context) error {
	for {
		select {
		case <-h.ticker.C:
			h.logger.Info("Hello!")
		case <-ctx.Done():
			return nil
		}
	}
}
