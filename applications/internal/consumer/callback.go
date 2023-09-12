package consumer

import (
	"context"
	"fmt"

	"codetalks/internal"
	"github.com/justtrackio/gosoline/pkg/cfg"
	"github.com/justtrackio/gosoline/pkg/ddb"
	"github.com/justtrackio/gosoline/pkg/log"
	"github.com/justtrackio/gosoline/pkg/mdl"
	"github.com/justtrackio/gosoline/pkg/stream"
)

var DdbTripRepositorySettings = &ddb.Settings{
	ModelId: mdl.ModelId{
		Name: "trips",
	},
	Main: ddb.MainSettings{
		Model: internal.Trip{},
	},
}

var DdbVendorRepositorySettings = &ddb.Settings{
	ModelId: mdl.ModelId{
		Name: "vendors",
	},
	Main: ddb.MainSettings{
		Model: Vendor{},
	},
}

type Callback struct {
	logger           log.Logger
	tripRepository   ddb.Repository
	vendorRepository ddb.Repository
}

func NewCallback(ctx context.Context, config cfg.Config, logger log.Logger) (stream.ConsumerCallback, error) {
	var err error
	var tripRepository, vendorRepository ddb.Repository

	if tripRepository, err = ddb.NewRepository(ctx, config, logger, DdbTripRepositorySettings); err != nil {
		return nil, fmt.Errorf("can not create trips repository: %w", err)
	}

	if vendorRepository, err = ddb.NewRepository(ctx, config, logger, DdbVendorRepositorySettings); err != nil {
		return nil, fmt.Errorf("can not create trips repository: %w", err)
	}

	consumer := &Callback{
		logger:           logger,
		tripRepository:   tripRepository,
		vendorRepository: vendorRepository,
	}

	return consumer, nil
}

func (c Callback) GetModel(attributes map[string]interface{}) interface{} {
	return &internal.Trip{}
}

func (c Callback) Consume(ctx context.Context, model interface{}, attributes map[string]interface{}) (bool, error) {
	var err error
	var result *ddb.GetItemResult

	logger := c.logger.WithContext(ctx)
	trip := model.(*internal.Trip)

	vendor := &Vendor{}
	qb := c.vendorRepository.GetItemBuilder().WithHash(trip.VendorID)

	if result, err = c.vendorRepository.GetItem(ctx, qb, vendor); err != nil {
		return false, fmt.Errorf("can not get vendor with id %d: %w", trip.VendorID, err)
	}

	if !result.IsFound {
		logger.Warn("there was no vendor found with id %d", trip.VendorID)
		return true, nil
	}

	if _, err = c.tripRepository.PutItem(ctx, nil, trip); err != nil {
		return false, fmt.Errorf("can not put trip: %w", err)
	}

	c.logger.WithContext(ctx).Info("stored trip from vendor %s", vendor.Name)

	return true, nil
}
