package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strings"
	"time"
	"encoding/json"
	"math/rand"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func gpsRegisterHandler(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("error upgrading GET request to a websocket:", err)
	}
	defer conn.Close()
	db, err := sql.Open("mysql", dbPath)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	event_id := rand.Intn(256*256*256-1)
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("error websocket reading:", err)
			return
		}
		var position GpsPosition
		err = json.Unmarshal(p, &position)
		if err != nil {
			log.Println("parse json error:", err)
			return
		}
		ok := position.Validate()
		if !ok {
			log.Println("some error occured. Validation failed.")
			return
		}
		fmt.Println(position)
		date_time := time.Now()
		date_split := strings.Split(date_time.String(), " ")
		date := date_split[0] + " " + date_split[1]

		var query string = fmt.Sprintf("INSERT INTO gps (event_id, date, latitude, longitude, altitude, accuracy, altitudeAccuracy, heading, speed) VALUES (%d, '%s', %g, %g, %g, %g, %g, %g, %g)", event_id, date, position.Latitude, position.Longitude, position.Altitude, position.Accuracy, position.AltitudeAccuracy, position.Heading, position.Speed)
		log.Println(query)
		_, err = db.Exec(query)
		if err != nil {
			log.Println("err writing db:", err)
		}
	}
	db.Close()
}
