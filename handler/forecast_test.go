package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hilton-james/forecast/config"
	"github.com/hilton-james/forecast/utils"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestGetForecast(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	config := config.Config{
		Forecast: config.Forecast{
			ListenPort: ":8080",
		},
		Debug: true,
	}
	forecastHandler := NewForecast(config, logger)

	tests := []struct {
		name           string
		lat            string
		long           string
		expectedStatus int
	}{
		{
			name:           "true test",
			lat:            "31.7403",
			long:           "-83.6460",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "false test",
			lat:            "false",
			long:           "false",
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/forecast?lat="+tt.lat+"&long="+tt.long, nil)
			w := httptest.NewRecorder()

			handler := utils.HandleApiError(forecastHandler.GetForecast)
			handler.ServeHTTP(w, req)

			resp := w.Result()
			defer resp.Body.Close()
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)
		})
	}
}
