package gateway

import (
	"context"
	"fmt"
	"net/http"

	"codetalks/internal"
	"github.com/justtrackio/gosoline/pkg/apiserver"
	"github.com/justtrackio/gosoline/pkg/cfg"
	"github.com/justtrackio/gosoline/pkg/log"
	"github.com/justtrackio/gosoline/pkg/stream"
)

type TripHandler struct {
	logger   log.Logger
	producer stream.Producer
}

func NewTripHandler(ctx context.Context, config cfg.Config, logger log.Logger) (*TripHandler, error) {
	var err error
	var producer stream.Producer

	if producer, err = stream.NewProducer(ctx, config, logger, "trips"); err != nil {
		return nil, fmt.Errorf("can not create trips producer: %w", err)
	}

	handler := &TripHandler{
		logger:   logger.WithChannel("handler"),
		producer: producer,
	}

	return handler, nil
}

func (t TripHandler) GetInput() interface{} {
	return &internal.Trip{}
}

func (t TripHandler) Handle(ctx context.Context, request *apiserver.Request) (*apiserver.Response, error) {
	trip := request.Body.(*internal.Trip)

	ctx = log.AppendLoggerContextField(ctx, log.Fields{
		"trip_uuid": trip.UUID,
	})

	if err := t.producer.WriteOne(ctx, trip); err != nil {
		return nil, fmt.Errorf("can not write trip: %w", err)
	}

	t.logger.WithContext(ctx).Info("wrote trip from vendor with id %d", trip.VendorID)

	return apiserver.NewStatusResponse(http.StatusOK), nil
}
