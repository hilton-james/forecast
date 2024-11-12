package external

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/hilton-james/forecast/model"
	"go.uber.org/zap"
)

var (
	ErrUnableToFetchForecastAddress = errors.New("unable to fetch foreCast address")
)

type Forecast struct {
	logger *zap.Logger
}

func NewForecast(logger *zap.Logger) *Forecast {
	return &Forecast{
		logger: logger,
	}
}

func (f *Forecast) FetchFromApi(lat, long string) (*model.ExternalForecast, error) {
	var (
		err             error
		forecastAddress struct {
			Address string `json:"forecast"`
		}
	)
	// curl -X GET "https://api.weather.gov/points/X,Y" -H "accept: application/ld+json"
	{
		var (
			body       []byte
			apiAddress string
		)

		apiAddress = fmt.Sprintf("https://api.weather.gov/points/%s,%s", lat, long)
		body, err = f.makeGetRequest(apiAddress)
		if err != nil {
			f.logger.Debug("api error", zap.Error(err))
			return nil, err
		}

		err = json.Unmarshal(body, &forecastAddress)
		if err != nil {
			f.logger.Debug("api unmarshal error", zap.Error(err))
			return nil, err
		}
		// f.logger.Info("fetched forecast Address", zap.String("address", forecastAddress.Address))

		//TODO: check more  validation factors
		if len(forecastAddress.Address) == 0 {
			f.logger.Debug("external api error", zap.Error(ErrUnableToFetchForecastAddress))
			return nil, ErrUnableToFetchForecastAddress
		}
	}

	//curl -X GET "https://api.weather.gov/gridpoints/HNX/X,Y/forecast" -H "accept: application/ld+json"
	var forecastInformation model.ExternalForecast
	{
		var body []byte

		body, err = f.makeGetRequest(forecastAddress.Address)
		if err != nil {
			f.logger.Debug("api error", zap.Error(err))
			return nil, err
		}
		// log.Printf("%s", string(body))

		err = json.Unmarshal(body, &forecastInformation)
		if err != nil {
			f.logger.Debug("api forecast information unmarshal error", zap.Error(err))
			return nil, err
		}
		// f.logger.Info("fetched forecast information", zap.Any("information", forecastInformation))

	}

	{
		//TODO: Check whether the fetched information is correct
	}
	return &forecastInformation, nil
}

func (f *Forecast) makeGetRequest(address string) ([]byte, error) {
	var (
		//TODO: implement retry if need be. find the reasonable amount for the timeout
		client   = &http.Client{Timeout: time.Duration(2000 * time.Millisecond)}
		request  *http.Request
		response *http.Response
		err      error
	)

	request, err = http.NewRequest("GET", address, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Accept", "application/ld+json")
	response, err = client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch:" + response.Status)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
