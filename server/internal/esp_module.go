package routes

import (
	"encoding/json"
	"fmt"

	"net/http"
)

// ESP8266 IP address (the IP you get after connecting the ESP to Wi-Fi)
const ESP8266_IP = "http://192.168.1.100" // Replace with your ESP IP address
// RelayState stores the state of the relay (whether it's on or off)
var relayState bool

// SensorData represents the motion detection or light sensor data from ESP8266
type SensorData struct {
	PhotoDiodeValue int `json:"photoDiodeValue"`
}

// LogMessage represents a log message received from the ESP8266 or actions on the server
type LogMessage struct {
	Log string `json:"log"`
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// Check the URL path to determine the action
	action := r.URL.Path[len("/"):]
	if action != "on" && action != "off" {
		http.Error(w, "Invalid action. Use /on or /off", http.StatusBadRequest)
		return
	}

	// Control the smart plug by calling the ESP8266
	err := controlSmartPlug(action)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send a success response
	fmt.Fprintf(w, "Smart plug turned %s", action)
}
func controlSmartPlug(action string) error {
	var url string

	// Set the URL based on the action
	if action == "on" {
		url = ESP8266_IP + "/on"
	} else if action == "off" {
		url = ESP8266_IP + "/off"
	} else {
		return fmt.Errorf("Invalid action: %s", action)
	}

	// Make an HTTP GET request to the ESP8266
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("Error sending request to ESP8266: %s", err)
	}
	defer resp.Body.Close()

	// Print the response (optional)
	fmt.Printf("Response: %s\n", resp.Status)
	return nil
}

// controlRelayHandler handles requests to control the relay
func controlRelayHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Read action from the request body (turn on/off)
		var requestBody struct {
			Action string `json:"action"`
		}
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&requestBody); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// If the action is "on", turn the relay on; "off" to turn it off
		if requestBody.Action == "on" {
			relayState = true
			controlRelay(true)
		} else if requestBody.Action == "off" {
			relayState = false
			controlRelay(false)
		} else {
			http.Error(w, "Invalid action", http.StatusBadRequest)
			return
		}

		// Respond with the current relay state
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("Relay state: %v", relayState)))
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

// controlRelay function simulates controlling the relay (you can replace it with actual relay control code)
func controlRelay(state bool) {
	if state {
		// Simulate turning on the relay (e.g., controlling the ESP8266 GPIO pin)
		fmt.Println("Relay ON")
		// Replace this with actual command to control the relay on ESP8266 via GPIO or other means
		// exec.Command("your-command-to-turn-on-relay").Run()
	} else {
		// Simulate turning off the relay
		fmt.Println("Relay OFF")
		// Replace this with actual command to control the relay on ESP8266 via GPIO or other means
		// exec.Command("your-command-to-turn-off-relay").Run()
	}
}

// motionDetectionHandler handles requests from the ESP8266 to send motion or light data
func motionDetectionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var data SensorData
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&data); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// If the photo diode value exceeds a threshold (indicating motion or light change)
		if data.PhotoDiodeValue > 500 {
			fmt.Println("Motion detected or light change detected!")
			// Trigger relay or any other action based on motion detection
			controlRelay(true) // Example: Turn on relay based on motion detection
		} else {
			fmt.Println("No motion or light change detected")
		}

		// Respond with status
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Sensor data received"))
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

// checkLightHandler checks light levels based on the photodiode value
func checkLightHandler(w http.ResponseWriter, r *http.Request) {
	// Implement logic here if you want to check light levels or handle light control
	// For example, you could adjust the relay state based on light levels
	fmt.Println("Checking light levels...")
	w.Write([]byte("Light check initiated"))
}
