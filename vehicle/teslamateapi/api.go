package teslamateapi

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/evcc-io/evcc/util"
	"github.com/evcc-io/evcc/util/request"
)

const (
	VehiclesURL = "cars"
	StatusURL   = "cars/%v/status"
)

// TODO Suuport fuer API-TOKEN fuer Provider:
// - MaxCurrent
// - StartCharge
// - StopCharge
// ErrAuthFail indicates authorization failure
var ErrAuthFail = errors.New("authorization failed")

// API implements teslamate api.
type API struct {
	*request.Helper
	baseURI string
}

type Requester interface {
	Request(*http.Request) error
}

// New creates a new teslamate API
func NewAPI(log *util.Logger, baseURI string) *API {
	v := &API{
		Helper:  request.NewHelper(log),
		baseURI: strings.TrimSuffix(baseURI, "/api/v1") + "/api/v1",
	}

	// TODO Suuport fuer API-TOKEN fuer Provider:
	// - MaxCurrent
	// - StartCharge
	// - StopCharge
	/*
		v.Client.Transport = &transport.Decorator{
			Base: v.Client.Transport,
			Decorator: transport.DecorateHeaders(map[string]string{
				"Authorization: ": Bearer <token>,
				"Content-Type":  "application/json",
			}),
		}
	*/
	return v
}

func (v *API) Vehicles() ([]Vehicle, error) {
	var res ApiCarsJSONData

	uri := fmt.Sprintf("%s/%s", v.baseURI, VehiclesURL)
	err := v.GetJSON(uri, &res)

	return res.Data.Cars, err
}

// StatusLatest retrieves the latest server-side status
func (v *API) Status(carid int) (ApiStatusData, error) {
	var res ApiStatusJSONData

	uri := fmt.Sprintf("%s/%s", v.baseURI, fmt.Sprintf(StatusURL, carid))
	err := v.GetJSON(uri, &res)

	return res.Data, err
}
