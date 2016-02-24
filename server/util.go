package server

import (
	"encoding/json"
	"net/http"
)

func writeJSON(w http.ResponseWriter, obj interface{}) {
	serialized, _ := json.MarshalIndent(obj, "", "  ")
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(serialized))
}
