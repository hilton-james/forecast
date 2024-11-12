package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/hilton-james/forecast/config"
	"github.com/hilton-james/forecast/request"
	"github.com/hilton-james/forecast/service"
	"go.uber.org/zap"
)

type Forecast struct {
	config  config.Config
	logger  *zap.Logger
	service *service.Forecast
}

func NewForecast(config config.Config, logger *zap.Logger) *Forecast {
	return &Forecast{
		config:  config,
		logger:  logger,
		service: service.NewForecast(logger.Named("service")),
	}
}

func (f Forecast) GetForecast(w http.ResponseWriter, r *http.Request) error {
	request := request.Forecast{
		Latitude:  r.URL.Query().Get("lat"),
		Longitude: r.URL.Query().Get("long"),
	}

	{
		if err := request.Valid(); err != nil {
			f.logger.Debug("invalid request error", zap.String("Longitude", request.Longitude), zap.String("Latitude", request.Latitude))
			return errors.New("invalid request")
		}
	}

	{
		response, err := f.service.FetchForecast(request.Latitude, request.Longitude)
		if err != nil {
			f.logger.Debug("service error", zap.Error(err))
			return errors.New("failed to fetch")
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
	return nil
}
