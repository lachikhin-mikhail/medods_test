package api

import (
	"encoding/json"
	"log"
	"net/http"
)

func writeErr(err error, w http.ResponseWriter) {
	log.Println(err)
	w.WriteHeader(http.StatusBadRequest)
	errResp := map[string]string{
		"error": err.Error(),
	}
	resp, err := json.Marshal(errResp)
	if err != nil {
		log.Println(err)
	}
	w.Write(resp)
}
