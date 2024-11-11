package main

import (
	"log"
	"net/http"

	"github.com/hilton-james/forecast/config"
	"github.com/hilton-james/forecast/handler"
	"go.uber.org/zap"
)

func main() {
	config := config.New()

	var logger *zap.Logger
	{
		var err error
		if config.Debug == true {
			logger, err = zap.NewDevelopment()
		} else {
			logger, err = zap.NewProduction()
		}
		if err != nil {
			log.Fatalf("zap creation error: %+v\n", err)
		}
	}
	logger.Info("application is running with this configuration", zap.Any("config", config))

	forecastServer := http.Server{}
	{
		forecastHTTP := http.NewServeMux()
		forecast := handler.NewForecast(config, logger.Named("forecast"))
		forecastServer = http.Server{
			Addr:    config.Forecast.ListenPort,
			Handler: forecastHTTP,
		}

		forecastHTTP.HandleFunc("GET /forecast", forecast.GetForecast)
	}
	if err := forecastServer.ListenAndServe(); err != http.ErrServerClosed {
		logger.Fatal("server error", zap.Error(err))
	}
}
