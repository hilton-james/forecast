package service

import (
	"github.com/hilton-james/forecast/external"
	"github.com/hilton-james/forecast/model"
	"go.uber.org/zap"
)

type Forecast struct {
	logger   *zap.Logger
	external external.Forecast
}

func NewForecast(logger *zap.Logger) *Forecast {
	return &Forecast{
		logger:   logger,
		external: *external.NewForecast(logger.Named("external")),
	}
}

func (f *Forecast) FetchForecast(lat, long string) (*model.Forecast, error) {
	var (
		err      error
		response model.Forecast
	)
	// f.logger.Info("service", zap.String("lat", lat), zap.String("long", long))

	var forecastInformation *model.ExternalForecast
	{
		forecastInformation, err = f.external.FetchFromApi(lat, long)
		if err != nil {
			f.logger.Debug("external api error", zap.Error(err))
			return nil, err
		}
	}

	{
		response.Forecast = forecastInformation.Periods[0].ShortForecast
		temperature := forecastInformation.Periods[0].Temperature
		unit := forecastInformation.Periods[0].TemperatureUnit
		response.Temperature = f.labelTemperature(temperature, unit)
	}
	return &response, nil
}

func (f *Forecast) labelTemperature(temperature int, unit string) string {

	if unit == "F" {
		temperature = (temperature - 32) * 5 / 9
	}

	if temperature <= 15 {
		return "cold"
	} else if temperature <= 25 {
		return "moderate"
	} else if temperature <= 35 {
		return "hot"
	}
	return "very hot"
}
