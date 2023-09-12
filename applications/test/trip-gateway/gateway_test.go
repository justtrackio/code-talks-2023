package trip_gateway

import (
	"net/http"
	"testing"

	"codetalks/internal"
	"codetalks/internal/gateway"
	"github.com/go-resty/resty/v2"
	"github.com/justtrackio/gosoline/pkg/apiserver"
	"github.com/justtrackio/gosoline/pkg/test/suite"
)

type GatewayTestSuite struct {
	suite.Suite
}

func TestApiServerTestSuite(t *testing.T) {
	suite.Run(t, new(GatewayTestSuite))
}

func (s *GatewayTestSuite) SetupSuite() []suite.Option {
	return []suite.Option{
		suite.WithConfigFile("../../cmd/trip-gateway/config.dist.yml"),
	}
}

func (s *GatewayTestSuite) SetupApiDefinitions() apiserver.Definer {
	return gateway.Definer
}

func (s *GatewayTestSuite) TestClick() *suite.ApiServerTestCase {
	return &suite.ApiServerTestCase{
		Method: http.MethodPost,
		Url:    "/trip",
		Body: &internal.Trip{
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
		},
		ExpectedStatusCode: http.StatusOK,
		Assert: func(response *resty.Response) error {
			output := s.Env().StreamOutput("trips")

			s.Equal(1, output.Len())

			trip := &internal.Trip{}
			output.Unmarshal(0, trip)

			s.Equal("a7b5ea8b-d19f-442d-8514-2339c04885a1", trip.UUID)

			return nil
		},
	}
}
