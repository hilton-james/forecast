package request

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type Forecast struct {
	Longitude string
	Latitude  string
}

func (f Forecast) Valid() error {
	return validation.ValidateStruct(&f,
		validation.Field(&f.Latitude, validation.Required, validation.Length(1, 10)),
		validation.Field(&f.Longitude, validation.Required, validation.Length(1, 12)),
	)
}
