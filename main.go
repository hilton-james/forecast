package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/hilton-james/forecast/config"
	"github.com/hilton-james/forecast/handler"
	"github.com/hilton-james/forecast/utils"
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

		forecastHTTP.HandleFunc("GET /forecast", utils.HandleApiError(forecast.GetForecast))
	}

	{
		stop := make(chan os.Signal)
		serverError := make(chan error)
		signal.Notify(stop, os.Kill, os.Interrupt)
		go func() {
			if err := forecastServer.ListenAndServe(); err != http.ErrServerClosed {
				serverError <- err
			}
		}()

		select {
		case err := <-serverError:
			log.Fatalf("server error: %+v\n", err)
		case <-stop:
			ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			defer cancel()
			if err := forecastServer.Shutdown(ctx); err != nil {
				log.Fatalf("shutdown error: %+v\n", err)
			}
			log.Println("closed")
		}
	}
}
