package consumer

import (
	"context"

	"github.com/justtrackio/gosoline/pkg/fixtures"
)

var Fixtures = []*fixtures.FixtureSet{
	{
		Enabled: true,
		Writer:  fixtures.DynamoDbFixtureWriterFactory(DdbVendorRepositorySettings),
		Fixtures: []any{
			&Vendor{
				Id:   2,
				Name: "NYC Yellow Cab Taxi",
			},
		},
	},
}

func FixtureBuilderFactory(ctx context.Context) fixtures.FixtureBuilder {
	return &FixtureBuilder{}
}

type FixtureBuilder struct {
}

func (b *FixtureBuilder) Fixtures() []*fixtures.FixtureSet {
	return []*fixtures.FixtureSet{
		{
			Enabled: true,
			Writer:  fixtures.DynamoDbFixtureWriterFactory(DdbVendorRepositorySettings),
			Fixtures: []any{
				&Vendor{
					Id:   2,
					Name: "NYC Yellow Cab Taxi",
				},
			},
		},
	}
}
