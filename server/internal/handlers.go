package routes

import (
	"encoding/json"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}
func registerCamera(w http.ResponseWriter, r *http.Request) {
	var cam Camera
	if err := json.NewDecoder(r.Body).Decode(&cam); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	mutex.Lock()
	streams[cam.StreamKey] = cam
	mutex.Unlock()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Camera registered"}`))
}

func availableStreams(w http.ResponseWriter, r *http.Request) {
	streamKeys := []string{}
	for key := range streams {
		streamKeys = append(streamKeys, key)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(streamKeys)
}
