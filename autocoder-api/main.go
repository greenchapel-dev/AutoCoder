package main

import (
	// "autocoder-api/obd_device"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rzetterberg/elmobd"
)

// Struct for holding all the data from the car and for returning as json
type CarData struct {
	RPM         string `json:"rpm"`
	CoolantTemp string `json:"coolantTemp"`
}

var obdDevice *elmobd.Device
var rpm *elmobd.EngineRPM
var carData CarData

// Create the API layer
func createAPI() {
	r := mux.NewRouter()
	r.HandleFunc("/get-rpm", GetRpmResp).Methods("GET")
	r.HandleFunc("/get-coolant-temp", GetCoolantTempResp).Methods("GET")

	// all data from the car
	r.HandleFunc("/all-data", GetCarData).Methods("GET")

	http.ListenAndServe(":8080", r)
}

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
	ticker := time.NewTicker(100 * time.Millisecond)
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
	ticker := time.NewTicker(1000 * time.Millisecond)
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

// Get the rpm API Response
func GetRpmResp(w http.ResponseWriter, r *http.Request) {
	rpmVal := GetRPM()
	// create response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(rpmVal))
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

// Get Coolant temp API Response
func GetCoolantTempResp(w http.ResponseWriter, r *http.Request) {
	coolentTempVal := GetCoolantTemp()
	// create response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(coolentTempVal))
}

// Get all the data from the car
func GetCarData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(carData)
}

// MAIN
func main() {
	addr := flag.String(
		"addr",
		// "test:///dev/ttyUSB0",
		"tcp://127.0.0.1:35000",
		"Address of the ELM327 device to use (use either test://, tcp://ip:port or serial:///dev/ttyS0)",
	)
	debug := flag.Bool(
		"debug",
		true,
		"Enable debug outputs",
	)

	flag.Parse()

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

	createAPI()
}