package main

// modified from https://github.com/rzetterberg/elmobd

// // Check the car has these commands
// func checkSupportedObdCommands() {
// 	supported, err := obdDevice.CheckSupportedCommands()

// 	if err != nil {
// 		fmt.Println("Failed to check supported commands", err)
// 		return
// 	}

// 	rpm = elmobd.NewEngineRPM()

// 	if supported.IsSupported(rpm) {
// 		fmt.Println("The car supports checking RPM")
// 	} else {
// 		fmt.Println("The car does NOT supports checking RPM")
// 	}
// }

// // Thread for getting the fast data from the car on a regular interval
// func getFastData() {
// 	// create a ticker for every 100ms
// 	ticker := time.NewTicker(100 * time.Millisecond)
// 	// run forever
// 	for {
// 		select {
// 		case <-ticker.C:
// 			// get the rpm
// 			GetRPM()
// 		}
// 	}
// }

// // Thread for getting the slow data from the car on a regular interval
// func getSlowData() {
// 	// create a ticker for every 1000ms
// 	ticker := time.NewTicker(1000 * time.Millisecond)
// 	// run forever
// 	for {
// 		select {
// 		case <-ticker.C:
// 			// get the rpm
// 			GetCoolantTemp()
// 		}
// 	}
// }

// // Get RPM
// func GetRPM() string {
// 	rpm = elmobd.NewEngineRPM()
// 	rmpRes, err := obdDevice.RunOBDCommand(rpm)
// 	if err != nil {
// 		fmt.Println("error", err)
// 		return "error: " + err.Error()
// 	}
// 	rpmVal := rmpRes.ValueAsLit()
// 	carData.RPM = rpmVal
// 	return rpmVal
// }

// // Get Engine coolent temp
// func GetCoolantTemp() string {
// 	coolantTemp := elmobd.NewCoolantTemperature()
// 	coolantTempRes, err := obdDevice.RunOBDCommand(coolantTemp)
// 	if err != nil {
// 		fmt.Println("error", err)
// 		return "error: " + err.Error()
// 	}
// 	coolantTempVal := coolantTempRes.ValueAsLit()
// 	carData.CoolantTemp = coolantTempVal
// 	return coolantTempVal
// }
