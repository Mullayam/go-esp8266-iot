package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	router := mux.NewRouter()
	router.Use(mux.CORSMethodMiddleware(router))

	// Register the handler for the routes
	router.HandleFunc("/", home).Methods(http.MethodGet)
	router.HandleFunc("/control-relay", controlRelayHandler).Methods(http.MethodPost)
	router.HandleFunc("/motion-detection", motionDetectionHandler).Methods(http.MethodPost)
	router.HandleFunc("/check-light", checkLightHandler).Methods(http.MethodPost)
	router.HandleFunc("/google-mini", handleRequest).Methods(http.MethodPost)

	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./static/"))))

	return router

}
