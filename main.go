package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"

	"github.com/joho/godotenv"
)

func init() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func main() {
	http.HandleFunc("/info", infoHandler) // Handle all routes
	log.Println("Server starting on http://localhost:8080...")
	http.ListenAndServe(":8080", nil)
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	// Extract IP and provider from query params
	queryParams := r.URL.Query()
	ip := queryParams.Get("ip")
	provider := queryParams.Get("provider")

	var ipProvider IPInfoProvider

	// Choose the provider based on the query parameter
	switch provider {
	case "ipinfo":
		ipProvider = ipinfoProvider{}
	default:
		ipProvider = ipAPIProvider{}
	}

	if ip == "" {
		ip, _, _ = net.SplitHostPort(r.RemoteAddr) // Default to the request's IP if none is provided
	}

	ipDetails, err := ipProvider.FetchIPInfo(ip)
	if err != nil {
		http.Error(w, "Failed to fetch IP information", http.StatusInternalServerError)
		log.Printf("Error fetching IP information: %v", err)
		return
	}

	prettyJSON, err := json.MarshalIndent(ipDetails, "", "  ")
	if err != nil {
		http.Error(w, "Failed to generate JSON", http.StatusInternalServerError)
		log.Printf("Error generating pretty JSON: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(prettyJSON)
}
