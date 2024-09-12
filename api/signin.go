package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lachikhin-mikhail/medods_test/internal/auth"
	"github.com/lachikhin-mikhail/medods_test/internal/db"
)

// PostSigninHandler обрабатывает запросы POST к /api/signin. Отправляет ответ с access и refresh токеном, или ошибку.
func PostSigninHandler(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	var body map[string]string
	var err error
	var guid string

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
	} else {
		guid = body["guid"]
	}
	credValid, err := db.VerifyUser(guid)
	if err != nil {
		writeErr(err, w)
		return
	}
	if credValid {
		access, err := auth.GenerateAccessToken(guid)
		if err != nil {
			writeErr(err, w)
			return
		}
		refresh, err := auth.GenerateRefreshToken(guid)
		if err != nil {
			writeErr(err, w)
			return
		}
		tokens := map[string]string{
			"access":  access, // change to call for token gen
			"refresh": refresh,
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
	var token, guid string

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
	} else {
		token = body["refresh"]
		guid = body["guid"]
	}
	if valid, err := db.VerifyRefreshToken(guid, token); valid && err == nil {
		access, err := auth.GenerateAccessToken(guid)
		if err != nil {
			writeErr(err, w)
			return
		}
		refresh, err := auth.GenerateRefreshToken(guid)
		if err != nil {
			writeErr(err, w)
			return
		}
		tokens := map[string]string{
			"access":  access,
			"refresh": refresh,
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
