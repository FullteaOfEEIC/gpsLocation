package main

import (
  "net/http"
  "log"
  "fmt"
  "encoding/json"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

func gpsShowHandler(w http.ResponseWriter, r *http.Request){
  db, err := sql.Open("mysql", dbPath)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
  var query string = "SELECT * FROM gps"
  rows, err := db.Query(query)
  if err != nil {
    log.Fatalln("DB load Failed",err)
  }
  defer rows.Close()
  var locations []GpsPositionWithEventID
  for rows.Next(){
    var location GpsPositionWithEventID

    var id int64
    var date string
    err := rows.Scan(&id, &location.Event_ID, &date, &location.Latitude, &location.Longitude, &location.Altitude,&location.Accuracy, &location.AltitudeAccuracy, &location.Heading, &location.Speed)
    if err != nil {
      log.Fatalln("Db parse Failed", err)
    }
    locations = append(locations, location)
  }
  e, err :=json.Marshal(locations)
  fmt.Fprintln(w, string(e))
  db.Close()
}
