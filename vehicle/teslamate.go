package vehicle

import (
	"time"

	"github.com/evcc-io/evcc/api"
	"github.com/evcc-io/evcc/util"
	"github.com/evcc-io/evcc/vehicle/teslamate"
)

// Teslamate is an api.Vehicle implementation for Tesla cars
type Teslamate struct {
	*embed
	*teslamate.Provider
}

func init() {
	registry.Add("teslamate", NewTeslamateFromConfig)
}

// NewTeslamateFromConfig creates a new vehicle
func NewTeslamateFromConfig(other map[string]interface{}) (api.Vehicle, error) {
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

	log := util.NewLogger("teslamate").Redact(cc.Token, cc.VIN)

	var err error

	// TODO Suuport fuer API-TOKEN fuer Provider:
	// - MaxCurrent
	// - StartCharge
	// - StopCharge
	api := teslamate.NewAPI(log, cc.ApiURI)

	vehicle, err := ensureVehicleEx(
		cc.VIN, api.Vehicles,
		func(v teslamate.Vehicle) string {
			return v.CarDetails.Vin
		},
	)

	if err != nil {
		return nil, err
	}

	v := &Teslamate{
		embed:    &cc.embed,
		Provider: teslamate.NewProvider(api, vehicle.CarID, cc.Cache),
	}

	if v.Title_ == "" {
		v.Title_ = vehicle.Name
	}

	return v, err
}
