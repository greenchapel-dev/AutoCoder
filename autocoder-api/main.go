package main

import (
	// "autocoder-api/obd_device"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/greenchapel-dev/elmobd" // modified from https://github.com/rzetterberg/elmobd
)

// Struct for holding all the data from the car and for returning as json
type CarData struct {
	RPM         string `json:"rpm"`
	CoolantTemp string `json:"coolantTemp"`
}

var obdDevice *elmobd.Device
var rpm *elmobd.EngineRPM
var carData CarData
var obdRunning = false
var addr *string
var debug *bool

// Create the API layer
func createAPI() {
	r := mux.NewRouter()
	// all data from the car
	r.HandleFunc("/all-car-data", GetCarData).Methods("GET")
	r.HandleFunc("/start-obd", PostStartOBD).Methods("POST")
	http.ListenAndServe(":8080", r)
}

// Get all the data from the car
func GetCarData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if obdRunning == false {
		// status not ok
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(carData)
	}
}

// function with post request to start the obd
func PostStartOBD(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if obdRunning == false {
		// start the obd
		obdRunning = true
		// Create a new device
		dev, err := elmobd.NewDevice(*addr, *debug)
		if err != nil {
			fmt.Println("error", err)
			return
		}
		obdDevice = dev

		// Create the API
		// checkSupportedObdCommands()
		go getFastData()
		go getSlowData()
		w.WriteHeader(http.StatusOK)
	} else {
		// obd already running
		w.WriteHeader(http.StatusNotFound)
	}
}

// MAIN
func main() {
	// command line flags
	addr = flag.String(
		"addr",
		// "test:///dev/ttyUSB0",
		"tcp://127.0.0.1:35000",
		"Address of the ELM327 device to use (use either test://, tcp://ip:port or serial:///dev/ttyS0)",
	)
	debug = flag.Bool(
		"debug",
		true,
		"Enable debug outputs",
	)
	flag.Parse()

	// start the API layer
	createAPI()
}
