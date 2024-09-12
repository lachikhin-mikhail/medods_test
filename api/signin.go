package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// PostSigninHandler обрабатывает запросы POST к /api/signin. Отправляет ответ с access и refresh токеном, или ошибку.
func PostSigninHandler(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	var body map[string]string
	var err error

	// Функция для ответа ошибкой в формате JSON
	_, err = buf.ReadFrom(r.Body)
	if err != nil {
		writeErr(err, w)
		return
	}

	if err = json.Unmarshal(buf.Bytes(), &body); err != nil {
		writeErr(err, w)
		return
	}

	if !(len(body["guid"]) > 0) {
		err = fmt.Errorf("didnt receive credentials")
		writeErr(err, w)
		return
	}
	if true { // change for request to DB to confirm credentials
		tokens := map[string]string{
			"access":  "token", // change to call for token gen
			"refresh": "token",
		}
		resp, err := json.Marshal(tokens)
		if err != nil {
			writeErr(err, w)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(resp)
	}

}

// PostRefreshHandler обрабатывает запросы POST к /api/refresh. Возвращает токены access и refresh, или ошибку.
func PostRefreshHandler(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	var body map[string]string
	var err error

	// Функция для ответа ошибкой в формате JSON
	_, err = buf.ReadFrom(r.Body)
	if err != nil {
		writeErr(err, w)
		return
	}

	if err = json.Unmarshal(buf.Bytes(), &body); err != nil {
		writeErr(err, w)
		return
	}

	if !(len(body["refresh"]) > 0) || !(len(body["guid"]) > 0) {
		err = fmt.Errorf("didnt receive credentials")
		writeErr(err, w)
		return
	}
	if true { // change for request to DB to confirm credentials
		tokens := map[string]string{
			"access":  "token", // change to call for token gen
			"refresh": "token",
		}
		resp, err := json.Marshal(tokens)
		if err != nil {
			writeErr(err, w)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(resp)
	}
}
