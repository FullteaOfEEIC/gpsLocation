package main

import (
  "gopkg.in/go-playground/validator.v9"
)

type GpsPosition struct {
	Latitude         float64 `validate:"numeric"`
	Longitude        float64 `validate:"numeric"`
	Altitude         float64 `validate:"numeric"`
	Accuracy         float64 `validate:"numeric"`
	AltitudeAccuracy float64 `validate:"numeric"`
	Heading          float64 `validate:"numeric"`
	Speed            float64 `validate:"numeric"`
}

type GpsPositionWithEventID struct{
  GpsPosition
  Event_ID float64
}

func (form *GpsPosition) Validate() (ok bool) {
	err := validator.New().Struct(*form)
	return err == nil
}

const dbPath string = "root:root@tcp(gps_db:3306)/location"
