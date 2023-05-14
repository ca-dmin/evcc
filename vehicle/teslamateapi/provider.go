package teslamateapi

import (
	"time"

	"github.com/evcc-io/evcc/api"
	"github.com/evcc-io/evcc/provider"
)

// Provider implements the vehicle api
type Provider struct {
	statusG func() (ApiStatusData, error)
}

// NewProvider creates a vehicle api provider
func NewProvider(api *API, carid int, cache time.Duration) *Provider {
	impl := &Provider{
		statusG: provider.Cached(func() (ApiStatusData, error) {
			return api.Status(carid)
		}, cache),
	}
	return impl
}

var _ api.Battery = (*Provider)(nil)

// Soc implements the api.Vehicle interface
func (v *Provider) Soc() (float64, error) {
	res, err := v.statusG()

	if err == nil {
		return float64(res.MQTTInformation.BatteryDetails.BatteryLevel), nil
	}
	return 0, err
}

var _ api.ChargeState = (*Provider)(nil)

// Status implements the api.ChargeState interface
func (v *Provider) Status() (api.ChargeStatus, error) {
	status := api.StatusNone // disconnected

	res, err := v.statusG()

	if err == nil {
		if res.MQTTInformation.State == "charging" {
			return api.StatusC, nil
		}
		if res.MQTTInformation.ChargingDetails.PluggedIn {
			return api.StatusB, nil
		} else {
			return api.StatusA, nil
		}
	}

	return status, err
}

var _ api.ChargeRater = (*Provider)(nil)

// ChargedEnergy implements the api.ChargeRater interface
func (v *Provider) ChargedEnergy() (float64, error) {
	res, err := v.statusG()

	if err == nil {
		return res.MQTTInformation.ChargingDetails.ChargeEnergyAdded, nil
	}

	return 0, err
}

var _ api.VehicleRange = (*Provider)(nil)

// Range implements the api.VehicleRange interface
func (v *Provider) Range() (int64, error) {
	res, err := v.statusG()

	if err == nil {
		// XXX evtl. besser EstRange?
		// Dann aber andes als in tesla.go
		return int64(res.MQTTInformation.BatteryDetails.RatedBatteryRange), err
	}

	return 0, err
}

var _ api.VehicleOdometer = (*Provider)(nil)

// Odometer implements the api.VehicleOdometer interface
func (v *Provider) Odometer() (float64, error) {
	res, err := v.statusG()
	if err == nil {
		return res.MQTTInformation.Odometer, nil
	}

	return 0, err
}

var _ api.VehicleFinishTimer = (*Provider)(nil)

// FinishTime implements the api.VehicleFinishTimer interface
func (v *Provider) FinishTime() (time.Time, error) {
	res, err := v.statusG()

	if err == nil {
		t := time.Now()
		return t.Add(time.Duration(res.MQTTInformation.ChargingDetails.TimeToFullCharge) * time.Minute), err
	}

	return time.Time{}, err
}

var _ api.VehiclePosition = (*Provider)(nil)

// Position implements the api.VehiclePosition interface
func (v *Provider) Position() (float64, float64, error) {
	res, err := v.statusG()
	if err == nil {
		return res.MQTTInformation.CarGeodata.Latitude, res.MQTTInformation.CarGeodata.Longitude, nil
	}

	return 0, 0, err
}

var _ api.SocLimiter = (*Provider)(nil)

// TargetSoc implements the api.SocLimiter interface
func (v *Provider) TargetSoc() (float64, error) {
	res, err := v.statusG()
	if err == nil {
		return float64(res.MQTTInformation.ChargingDetails.ChargeLimitSoc), nil
	}

	return 0, err
}

/*
	// TODO Suuport fuer API-TOKEN fuer Provider:
	// - MaxCurrent
	// - StartCharge
	// - StopCharge


var _ api.CurrentLimiter = (*Provider)(nil)

func (v *Provider) MaxCurrent(current int64) error {
	return v.vehicle.SetChargingAmps(int(current))
}


var _ api.VehicleChargeController = (*Provider)(nil)

// StartCharge implements the api.VehicleChargeController interface
func (v *Provider) StartCharge() error {
	err := v.vehicle.StartCharging()

	if err != nil && err.Error() == "408 Request Timeout" {
		if _, err := v.vehicle.Wakeup(); err != nil {
			return err
		}

		timer := time.NewTimer(90 * time.Second)

		for {
			select {
			case <-timer.C:
				return api.ErrTimeout
			default:
				time.Sleep(2 * time.Second)
				if err := v.vehicle.StartCharging(); err == nil || err.Error() != "408 Request Timeout" {
					return err
				}
			}
		}
	}

	return err
}

// StopCharge implements the api.VehicleChargeController interface
func (v *Provider) StopCharge() error {
	err := v.vehicle.StopCharging()

	// ignore sleeping vehicle
	if err != nil && err.Error() == "408 Request Timeout" {
		err = nil
	}

	return err
}
*/
