package vehicle

import (
	"time"

	"github.com/evcc-io/evcc/api"
	"github.com/evcc-io/evcc/util"
	"github.com/evcc-io/evcc/vehicle/teslamateapi"
)

// Teslamateapi is an api.Vehicle implementation for Tesla cars
type Teslamateapi struct {
	*embed
	*teslamateapi.Provider
}

func init() {
	registry.Add("teslamateapi", NewTeslamateapiFromConfig)
}

// NewTeslamateapiFromConfig creates a new vehicle
func NewTeslamateapiFromConfig(other map[string]interface{}) (api.Vehicle, error) {
	cc := struct {
		embed              `mapstructure:",squash"`
		Token, VIN, ApiURI string
		Cache              time.Duration
	}{
		Cache: interval,
	}

	if err := util.DecodeOther(other, &cc); err != nil {
		return nil, err
	}

	log := util.NewLogger("teslamateapi").Redact(cc.Token, cc.VIN)

	var err error

	// TODO Suuport fuer API-TOKEN fuer Provider:
	// - MaxCurrent
	// - StartCharge
	// - StopCharge
	api := teslamateapi.NewAPI(log, cc.ApiURI)

	vehicle, err := ensureVehicleEx(
		cc.VIN, api.Vehicles,
		func(v teslamateapi.Vehicle) string {
			return v.CarDetails.Vin
		},
	)

	if err != nil {
		return nil, err
	}

	v := &Teslamateapi{
		embed:    &cc.embed,
		Provider: teslamateapi.NewProvider(api, vehicle.CarID, cc.Cache),
	}

	if v.Title_ == "" {
		v.Title_ = vehicle.Name
	}

	return v, err
}
