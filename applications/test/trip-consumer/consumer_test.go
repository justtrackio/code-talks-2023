//go:build integration && fixtures

package trip_consumer

import (
	"context"
	"testing"

	"codetalks/internal"
	"codetalks/internal/consumer"
	"github.com/justtrackio/gosoline/pkg/test/suite"
)

type ConsumerTestSuite struct {
	suite.Suite
}

func TestConsumerTestSuite(t *testing.T) {
	suite.Run(t, new(ConsumerTestSuite))
}

func (s *ConsumerTestSuite) SetupSuite() []suite.Option {
	return []suite.Option{
		suite.WithLogLevel("debug"),
		suite.WithConfigFile("../../cmd/trip-consumer/config.dist.yml"),
		suite.WithConsumer(consumer.NewCallback),
		suite.WithFixtureBuilderFactories(consumer.FixtureBuilderFactory),
	}
}

func (s *ConsumerTestSuite) TestSuccess() *suite.StreamTestCase {
	expected := &internal.Trip{
		UUID:                 "a7b5ea8b-d19f-442d-8514-2339c04885a1",
		VendorID:             2,
		TpepPickupDatetime:   "2023-03-01 00:06:43",
		TpepDropoffDatetime:  "2023-03-01 00:16:43",
		PassengerCount:       1,
		TripDistance:         0,
		RatecodeID:           1,
		StoreAndFwdFlag:      "N",
		PULocationID:         238,
		DOLocationID:         42,
		PaymentType:          2,
		FareAmount:           8.6,
		Extra:                1.0,
		MtaTax:               0.5,
		TipAmount:            0.0,
		TollsAmount:          0.0,
		ImprovementSurcharge: 1.0,
		TotalAmount:          11.1,
		CongestionSurcharge:  0,
		AirportFee:           0,
	}

	return &suite.StreamTestCase{
		Input: map[string][]suite.StreamTestCaseInput{
			"consumer": {
				{
					Body: expected,
				},
			},
		},
		Assert: func() error {
			repo, err := s.Env().DynamoDb("default").Repository(consumer.DdbTripRepositorySettings)
			s.NoError(err, "there should beno error on creating the repository")

			trip := &internal.Trip{}
			qb := repo.GetItemBuilder().WithHash("a7b5ea8b-d19f-442d-8514-2339c04885a1")
			_, err = repo.GetItem(context.Background(), qb, trip)

			s.NoError(err, "there should be no error on getting the trip")
			s.Equal(expected, trip)

			return nil
		},
	}
}
