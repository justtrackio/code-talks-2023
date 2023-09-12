package gateway

import (
	"context"
	"fmt"

	"github.com/justtrackio/gosoline/pkg/apiserver"
	"github.com/justtrackio/gosoline/pkg/cfg"
	"github.com/justtrackio/gosoline/pkg/log"
)

var Definer = func(ctx context.Context, config cfg.Config, logger log.Logger) (*apiserver.Definitions, error) {
	def := &apiserver.Definitions{}

	var err error
	var handler apiserver.HandlerWithInput

	if handler, err = NewTripHandler(ctx, config, logger); err != nil {
		return nil, fmt.Errorf("can not create trip handler: %w", err)
	}

	def.POST("/trip", apiserver.CreateJsonHandler(handler))

	return def, nil
}
