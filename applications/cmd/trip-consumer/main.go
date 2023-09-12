package main

import (
	"codetalks/internal/consumer"
	"github.com/justtrackio/gosoline/pkg/application"
	"github.com/justtrackio/gosoline/pkg/fixtures"
)

func main() {
	application.RunConsumer(
		consumer.NewCallback,
		application.WithFixtureBuilderFactory(fixtures.SimpleFixtureBuilderFactory(consumer.Fixtures)),
	)
}
