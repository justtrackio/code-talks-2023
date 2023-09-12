package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"time"

	"codetalks/internal"
	"github.com/justtrackio/gosoline/pkg/cfg"
	"github.com/justtrackio/gosoline/pkg/clock"
	"github.com/justtrackio/gosoline/pkg/coffin"
	gosoHttp "github.com/justtrackio/gosoline/pkg/http"
	"github.com/justtrackio/gosoline/pkg/kernel"
	"github.com/justtrackio/gosoline/pkg/log"
)

type ForwarderSettings struct {
	GatewayUrl string `cfg:"gateway_url"`
	Path       string `cfg:"path"`
}

type Forwarder struct {
	logger        log.Logger
	forwardTicker clock.Ticker
	freqTicker    clock.Ticker
	logTicker     clock.Ticker
	client        gosoHttp.Client
	file          *os.File
	reader        *bufio.Reader
	gatewayUrl    string
}

func NewForwarder(ctx context.Context, config cfg.Config, logger log.Logger) (kernel.Module, error) {
	var err error
	var client gosoHttp.Client
	var file *os.File

	settings := &ForwarderSettings{}
	config.UnmarshalKey("forwarder", settings)

	if client, err = gosoHttp.ProvideHttpClient(ctx, config, logger, "trips"); err != nil {
		return nil, fmt.Errorf("can not create http client: %w", err)
	}

	if file, err = os.Open(settings.Path); err != nil {
		return nil, fmt.Errorf("can not open trip file: %w", err)
	}

	module := &Forwarder{
		logger:        logger.WithChannel("forwarder"),
		forwardTicker: clock.NewRealTicker(time.Millisecond * 500),
		freqTicker:    clock.NewRealTicker(time.Minute),
		logTicker:     clock.NewRealTicker(time.Second * 10),
		client:        client,
		file:          file,
		reader:        bufio.NewReader(file),
		gatewayUrl:    settings.GatewayUrl,
	}

	return module, nil
}

func (f Forwarder) Run(ctx context.Context) error {
	c := 0

	cfn := coffin.New()
	cfn.GoWithContext(ctx, func(ctx context.Context) error {
		i := 0.0

		for {
			select {
			case <-f.freqTicker.Chan():
				dur := time.Duration((math.Abs(math.Sin(i*math.Pi/180))+1)*500) * time.Millisecond
				f.logger.Info("changing forward interval to %s", dur)
				f.forwardTicker.Reset(dur)
				i += 10
			case <-ctx.Done():
				return nil
			}
		}
	})
	cfn.GoWithContext(ctx, func(ctx context.Context) error {
		for {
			select {
			case <-f.forwardTicker.Chan():
				if err := f.forward(ctx); err != nil {
					f.logger.Error("can not forward trip: %s", err)
				}
				c++
			case <-ctx.Done():
				return nil
			}
		}
	})
	cfn.GoWithContext(ctx, func(ctx context.Context) error {
		for {
			select {
			case <-f.logTicker.Chan():
				f.logger.Info("forwarded %d trips", c)
				c = 0
			case <-ctx.Done():
				return nil
			}
		}
	})

	return cfn.Wait()
}

func (f Forwarder) forward(ctx context.Context) error {
	var err error
	var trip *internal.Trip
	var resp *gosoHttp.Response

	if trip, err = f.readTrip(); err != nil {
		return fmt.Errorf("can not read next trip: %w", err)
	}

	req := f.client.NewJsonRequest().
		WithBody(trip).
		WithUrl(f.gatewayUrl)

	if resp, err = f.client.Post(ctx, req); err != nil {
		return fmt.Errorf("can not get trip data from server: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-ok status")
	}

	return nil
}

func (f Forwarder) readTrip() (*internal.Trip, error) {
	var err error
	var data []byte

	for {
		if data, _, err = f.reader.ReadLine(); err != nil && err != io.EOF {
			return nil, fmt.Errorf("can not read line: %w", err)
		}

		if err == nil {
			break
		}

		if _, err = f.file.Seek(0, 0); err != nil {
			return nil, fmt.Errorf("can not reset trips file to the beginning: %w", err)
		}

		f.logger.Info("reached end of file: seeking to the beginning")
	}

	trip := &internal.Trip{}
	if err = json.Unmarshal(data, trip); err != nil {
		return nil, fmt.Errorf("can not unmarshal trip from json: %w", err)
	}

	return trip, nil
}
