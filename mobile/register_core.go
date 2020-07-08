package main

import (
	"database/sql"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
	"log"
	"net/http"
	"strconv"
	"time"
  "strings"
  _ "github.com/go-sql-driver/mysql"
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

func (form *GpsPosition) Validate() (ok bool) {
	err := validator.New().Struct(*form)
	return err == nil
}

func gpsRegisterHandler(w http.ResponseWriter, r *http.Request) {
	latitude_str := r.FormValue("latitude")
	latitude, err := strconv.ParseFloat(latitude_str, 64)
	if err != nil {
		latitude = 0
	}
	longitude_str := r.FormValue("longitude")
	longitude, err := strconv.ParseFloat(longitude_str, 64)
	if err != nil {
		longitude = 0
	}
	altitude_str := r.FormValue("altitude")
	altitude, err := strconv.ParseFloat(altitude_str, 64)
	if err != nil {
		altitude = 0
	}
	accuracy_str := r.FormValue("accuracy")
	accuracy, err := strconv.ParseFloat(accuracy_str, 64)
	if err != nil {
		accuracy = 0
	}

	altitudeAccuracy_str := r.FormValue("altitudeAccuracy")
	altitudeAccuracy, err := strconv.ParseFloat(altitudeAccuracy_str, 64)
	if err != nil {
		altitudeAccuracy = 0
	}
	heading_str := r.FormValue("heading")
	heading, err := strconv.ParseFloat(heading_str, 64)
	if err != nil {
		heading = 0
	}
	speed_str := r.FormValue("speed")
	speed, err := strconv.ParseFloat(speed_str, 64)
	if err != nil {
		speed = 0
	}
	position := GpsPosition{
		Latitude:         latitude,
		Longitude:        longitude,
		Altitude:         altitude,
		Accuracy:         accuracy,
		AltitudeAccuracy: altitudeAccuracy,
		Heading:          heading,
		Speed:            speed,
	}
	event_id := 0
	date_time := time.Now()
  date_split := strings.Split(date_time.String(), " ")
  date := date_split[0]+" "+date_split[1]

	ok := position.Validate()
	if !ok {
		log.Fatalln("some error occured")
	}

	db, err := sql.Open("mysql", "root:root@tcp(gps_db:3306)/location")
	if err != nil {
		log.Fatalln("DB access Failed", err)
	}
	defer db.Close()
	var query string = fmt.Sprintf("INSERT INTO gps (event_id, date, latitude, longitude, altitude, accuracy, altitudeAccuracy, heading, speed) VALUES (%d, '%s', %g, %g, %g, %g, %g, %g, %g)", event_id, date, position.Latitude, position.Longitude, position.Altitude, position.Accuracy, position.AltitudeAccuracy, position.Heading, position.Speed)

	log.Println(query)
  log.Println(date)
	_, err = db.Query(query)
	if err != nil {
		log.Fatalln("err",err)
	}

}