package teslamateapi

// from  https://github.com/tobiasehlert/teslamateapi/blob/main/src/v1_TeslaMateAPICars.go
// changes:
// - prefixed with ApiCars
// - ApiCarsData --> Vehicles - as it's evcc wording
// - ApiCarsCars --> Vehicle - as it's evcc wording

// JSONData struct - main
type ApiCarsJSONData struct {
	Data Vehicles `json:"data"`
}

// Information struct - child of JSONData
// ApiCarsData --> Vehicles
type Vehicles struct {
	Cars []Vehicle `json:"cars"`
}

// Cars struct - child of Data
// ApiCarsCars --> Vehicle
type Vehicle struct {
	//type  struct {
	CarID            int                     `json:"car_id"`            // smallint
	Name             string                  `json:"name"`              // text
	CarDetails       ApiCarsCarDetails       `json:"car_details"`       // struct
	Exterior         ApiCarsCarExterior      `json:"car_exterior"`      // struct
	Settings         ApiCarsCarSettings      `json:"car_settings"`      // struct
	TeslaMateDetails ApiCarsTeslaMateDetails `json:"teslamate_details"` // struct
	TeslaMateStats   ApiCarsTeslaMateStats   `json:"teslamate_stats"`   // struct
}

// creating structs for /cars
// CarDetails struct - child of Cars
type ApiCarsCarDetails struct {
	EID   int64  `json:"eid"`   // bigint
	VID   int64  `json:"vid"`   // bigint
	Vin   string `json:"vin"`   // text
	Model string `json:"model"` // character varying(255)
	//TrimBadging NullString  `json:"trim_badging"` // text
	//Efficiency  NullFloat64 `json:"efficiency"`   // double precision
}

// CarExterior struct - child of Cars
type ApiCarsCarExterior struct {
	ExteriorColor string `json:"exterior_color"` // text
	SpoilerType   string `json:"spoiler_type"`   // text
	WheelType     string `json:"wheel_type"`     // text
}

// CarSettings struct - child of Cars
type ApiCarsCarSettings struct {
	SuspendMin          int  `json:"suspend_min"`            // int
	SuspendAfterIdleMin int  `json:"suspend_after_idle_min"` // int
	ReqNotUnlocked      bool `json:"req_not_unlocked"`       // bool
	FreeSupercharging   bool `json:"free_supercharging"`     // bool
	UseStreamingAPI     bool `json:"use_streaming_api"`      // bool
}

// TeslaMateDetails struct - child of Cars
type ApiCarsTeslaMateDetails struct {
	InsertedAt string `json:"inserted_at"` // timestamp(0) without time zone
	UpdatedAt  string `json:"updated_at"`  // timestamp(0) without time zone
}

// TeslaMateStats struct - child of Cars
type ApiCarsTeslaMateStats struct {
	TotalCharges int `json:"total_charges"` // int
	TotalDrives  int `json:"total_drives"`  // int
	TotalUpdates int `json:"total_updates"` // int
}

// https://github.com/tobiasehlert/teslamateapi/blob/main/src/v1_TeslaMateAPICarsStatus.go
// changes:
// - prefixed with ApiStatus

// JSONData struct - main
type ApiStatusJSONData struct {
	Data ApiStatusData `json:"data"`
}

// Data struct - child of JSONData
type ApiStatusData struct {
	Car             ApiStatusCar             `json:"car"`
	MQTTInformation ApiStatusMQTTInformation `json:"status"`
	TeslaMateUnits  ApiStatusTeslaMateUnits  `json:"units"`
}

// MQTTInformation struct - child of Cars
type ApiStatusMQTTInformation struct {
	DisplayName     string                   `json:"display_name"`     // Blue Thunder - Vehicle Name
	State           string                   `json:"state"`            // asleep - Status of the vehicle (e.g. online, asleep, charging)
	StateSince      string                   `json:"state_since"`      // 2019-02-29T23:00:07Z - Date of the last status change
	Odometer        float64                  `json:"odometer"`         // 1653 - Car odometer in km
	CarStatus       ApiStatusCarStatus       `json:"car_status"`       // struct
	CarDetails      ApiStatusCarDetails      `json:"car_details"`      // struct
	CarExterior     ApiStatusCarExterior     `json:"car_exterior"`     // struct
	CarGeodata      ApiStatusCarGeodata      `json:"car_geodata"`      // struct
	CarVersions     ApiStatusCarVersions     `json:"car_versions"`     // struct
	DrivingDetails  ApiStatusDrivingDetails  `json:"driving_details"`  // struct
	ClimateDetails  ApiStatusClimateDetails  `json:"climate_details"`  // struct
	BatteryDetails  ApiStatusBatteryDetails  `json:"battery_details"`  // struct
	ChargingDetails ApiStatusChargingDetails `json:"charging_details"` // struct
	TpmsDetails     ApiStatusTpmsDetails     `json:"tpms_details"`     // struct
}

// Cars struct - child of Data
type ApiStatusCar struct {
	CarID   int    `json:"car_id"`   // smallint
	CarName string `json:"car_name"` // text
}

// TeslaMateUnits struct - child of Data
type ApiStatusTeslaMateUnits struct {
	UnitsLength      string `json:"unit_of_length"`      // string
	UnitsTemperature string `json:"unit_of_temperature"` // string
}

// creating structs for /cars
// BatteryDetails struct - child of MQTTInformation
type ApiStatusBatteryDetails struct {
	EstBatteryRange    float64 `json:"est_battery_range"`    // 372.5 - Estimated Range in km
	RatedBatteryRange  float64 `json:"rated_battery_range"`  // 401.63 - Rated Range in km
	IdealBatteryRange  float64 `json:"ideal_battery_range"`  // 335.79 - Ideal Range in km
	BatteryLevel       int     `json:"battery_level"`        // 88 - Battery Level Percentage
	UsableBatteryLevel int     `json:"usable_battery_level"` // 85 - Usable battery level percentage
}

// CarDetails struct - child of MQTTInformation
type ApiStatusCarDetails struct {
	Model       string `json:"model"`        // character varying(255)
	TrimBadging string `json:"trim_badging"` // P100D - Trim badging
}

// CarExterior struct - child of MQTTInformation
type ApiStatusCarExterior struct {
	ExteriorColor string `json:"exterior_color"` // DeepBlue - The exterior color
	SpoilerType   string `json:"spoiler_type"`   // None - The spoiler type
	WheelType     string `json:"wheel_type"`     // Pinwheel18 - The wheel type
}

// CarGeodata struct - child of MQTTInformation
type ApiStatusCarGeodata struct {
	Geofence  string  `json:"geofence"`  // Home - The name of the Geo-fence, if one exists at the current position
	Latitude  float64 `json:"latitude"`  // 35.278131 - Last reported car latitude
	Longitude float64 `json:"longitude"` // 29.744801 - Last reported car longitude
}

// CarStatus struct - child of MQTTInformation
type ApiStatusCarStatus struct {
	Healthy       bool `json:"healthy"`         // true - Health status of the logger for that vehicle
	Locked        bool `json:"locked"`          // true - Indicates if the car is locked
	SentryMode    bool `json:"sentry_mode"`     // false - Indicates if Sentry Mode is active
	WindowsOpen   bool `json:"windows_open"`    // false - Indicates if any of the windows are open
	DoorsOpen     bool `json:"doors_open"`      // false - Indicates if any of the doors are open
	TrunkOpen     bool `json:"trunk_open"`      // false - Indicates if the trunk is open
	FrunkOpen     bool `json:"frunk_open"`      // false - Indicates if the frunk is open
	IsUserPresent bool `json:"is_user_present"` // false - Indicates if a user is present in the vehicle
}

// CarVersions struct - child of MQTTInformation
type ApiStatusCarVersions struct {
	Version         string `json:"version"`          // 2019.32.12.2 - Software Version
	UpdateAvailable bool   `json:"update_available"` // false - Indicates if a software update is available
	UpdateVersion   string `json:"update_version"`   // 2019.32.12.3 - Software version of the available update
}

// ChargingDetails struct - child of MQTTInformation
type ApiStatusChargingDetails struct {
	PluggedIn                  bool    `json:"plugged_in"`                    // true - If car is currently plugged into a charger
	ChargeEnergyAdded          float64 `json:"charge_energy_added"`           // 5.06 - Last added energy in kWh
	ChargeLimitSoc             int     `json:"charge_limit_soc"`              // 90 - Charge Limit Configured in Percentage
	ChargePortDoorOpen         bool    `json:"charge_port_door_open"`         // true - Indicates if the charger door is open
	ChargerActualCurrent       float64 `json:"charger_actual_current"`        // 2.05 - Current amperage supplied by charger
	ChargerPhases              int     `json:"charger_phases"`                // 3 - Number of charger power phases (1-3)
	ChargerPower               float64 `json:"charger_power"`                 // 48.9 - Charger Power
	ChargerVoltage             int     `json:"charger_voltage"`               // 240 - Charger Voltage
	ChargeCurrentRequest       int     `json:"charge_current_request"`        // 40 - How many amps the car wants
	ChargeCurrentRequestMax    int     `json:"charge_current_request_max"`    // 40 - How many amps the car can have
	ScheduledChargingStartTime string  `json:"scheduled_charging_start_time"` // 2019-02-29T23:00:07Z - Start time of the scheduled charge
	TimeToFullCharge           float64 `json:"time_to_full_charge"`           // 1.83 - Hours remaining to full charge
}

// ClimateDetails struct - child of MQTTInformation
type ApiStatusClimateDetails struct {
	IsClimateOn       bool    `json:"is_climate_on"`      // true - Indicates if the climate control is on
	InsideTemp        float64 `json:"inside_temp"`        // 20.8 - Inside Temperature in °C
	OutsideTemp       float64 `json:"outside_temp"`       // 18.4 - Temperature in °C
	IsPreconditioning bool    `json:"is_preconditioning"` // false - Indicates if the vehicle is being preconditioned
}

// DrivingDetails struct - child of MQTTInformation
type ApiStatusDrivingDetails struct {
	ShiftState string `json:"shift_state"` // D - Current/Last Shift State (D/N/R/P)
	Power      int    `json:"power"`       // -9 Current battery power in watts. Positive value on discharge, negative value on charge
	Speed      int    `json:"speed"`       // 12 - Current Speed in km/h
	Heading    int    `json:"heading"`     // 340 - Last reported car direction
	Elevation  int    `json:"elevation"`   // 70 - Current elevation above sea level in meters
}

// TpmsDetails struct - child of MQTTInformatiojn
type ApiStatusTpmsDetails struct {
	TpmsPressureFL float64 `json:"tpms_pressure_fl"` // 2.9 - Tire pressure measure in BAR, front left tire
	TpmsPressureFR float64 `json:"tpms_pressure_fr"` // 2.8 - Tire pressure measure in BAR, front right tire
	TpmsPressureRL float64 `json:"tpms_pressure_rl"` // 2.9 - Tire pressure measure in BAR, rear left tire
	TpmsPressureRR float64 `json:"tpms_pressure_rr"` // 2.8 - Tire pressure measure in BAR, rear right tire
}
