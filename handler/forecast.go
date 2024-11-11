package handler

import (
	"net/http"

	"github.com/hilton-james/forecast/config"
	"go.uber.org/zap"
)

type Forecast struct {
	config config.Config
	logger *zap.Logger
}

func NewForecast(config config.Config, logger *zap.Logger) *Forecast {
	return &Forecast{
		config: config,
		logger: logger,
	}
}

func (f Forecast) GetForecast(w http.ResponseWriter, r *http.Request) {

	message := "hiii"
	w.Write([]byte(message))
}
