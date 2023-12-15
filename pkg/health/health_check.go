package health

import (
	"encoding/json"
	"log"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]bool)
	resp["alive"] = true
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Panicf("Error happened in JSON marshal. Err: %s", err)
	}
	_, err = w.Write(jsonResp)
	if err != nil {
		log.Panicf("Failed to check health")
	}
}
