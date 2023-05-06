package main

import (
	// "autocoder-api/obd_device"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"time"

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
		// "tcp://127.0.0.1:35000",
		"serial:///dev/ttyS3",
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

// / move to own file
// Check the car has these commands
func checkSupportedObdCommands() {
	supported, err := obdDevice.CheckSupportedCommands()

	if err != nil {
		fmt.Println("Failed to check supported commands", err)
		return
	}

	rpm = elmobd.NewEngineRPM()

	if supported.IsSupported(rpm) {
		fmt.Println("The car supports checking RPM")
	} else {
		fmt.Println("The car does NOT supports checking RPM")
	}
}

// Thread for getting the fast data from the car on a regular interval
func getFastData() {
	// create a ticker for every 100ms
	ticker := time.NewTicker(500 * time.Millisecond)
	// run forever
	for {
		select {
		case <-ticker.C:
			// get the rpm
			GetRPM()
		}
	}
}

// Thread for getting the slow data from the car on a regular interval
func getSlowData() {
	// create a ticker for every 1000ms
	ticker := time.NewTicker(2000 * time.Millisecond)
	// run forever
	for {
		select {
		case <-ticker.C:
			// get the rpm
			GetCoolantTemp()
		}
	}
}

// Get RPM
func GetRPM() string {
	rpm = elmobd.NewEngineRPM()
	rmpRes, err := obdDevice.RunOBDCommand(rpm)
	if err != nil {
		fmt.Println("error", err)
		return "error: " + err.Error()
	}
	rpmVal := rmpRes.ValueAsLit()
	carData.RPM = rpmVal
	return rpmVal
}

// Get Engine coolent temp
func GetCoolantTemp() string {
	coolantTemp := elmobd.NewCoolantTemperature()
	coolantTempRes, err := obdDevice.RunOBDCommand(coolantTemp)
	if err != nil {
		fmt.Println("error", err)
		return "error: " + err.Error()
	}
	coolantTempVal := coolantTempRes.ValueAsLit()
	carData.CoolantTemp = coolantTempVal
	return coolantTempVal
}
