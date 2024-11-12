package model

type Forecast struct {
	Forecast    string `json:"forecast"`
	Temperature string `json:"temperature"`
	// "forecast": "Partly Cloudy",
	// "temperature": "moderate"
}

type ExternalForecast struct {
	Periods []struct {
		Number          int    `json:"number"`
		Name            string `json:"name"`
		Temperature     int    `json:"temperature"`
		TemperatureUnit string `json:"temperatureUnit"`
		ShortForecast   string `json:"shortForecast"`
	} `json:"periods"`
}
