package utils

import (
	"encoding/json"
	"net/http"
)

func MessageData(status bool, message string, data interface{}) map[string]interface{} {
	resp := Message(status, message)
	resp["data"] = data
	return resp
}

func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func RespondError(w http.ResponseWriter, data map[string]interface{}, errcode int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(errcode)
	json.NewEncoder(w).Encode(data)
}
