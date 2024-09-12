package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	signinPath      = "api/signin"
	refreshPath     = "api/refresh"
	userGuidValid   = "ce547c40-acf9-11e6-80f5-76304dec7eb7"
	userGuidInvalid = "imaguidiswear:)"
)

// Sends request to /signin, recieves access/refresh tokens if user is correct
func TestSignin(t *testing.T) {
	godotenv.Load("../.env")
	var (
		m    map[string]any
		err  error
		user map[string]any
	)
	user = make(map[string]any)

	user["guid"] = userGuidValid

	body, err := request(signinPath, http.MethodPost, user)
	require.NoError(t, err)

	err = json.Unmarshal(body, &m)
	require.NoError(t, err)
	assert.NotEmpty(t, m["access"])
	assert.NotEmpty(t, m["refresh"])

}

// Sends request to /signin with false credentials, receives error
func TestSigninFalseCredentials(t *testing.T) {
	var (
		m    map[string]any
		user map[string]any
		err  error
	)
	user = make(map[string]any)
	user["guid"] = userGuidInvalid

	body, err := request(signinPath, http.MethodPost, user)
	require.NoError(t, err)
	err = json.Unmarshal(body, &m)
	require.NoError(t, err)
	assert.Empty(t, m["access"])
	assert.Empty(t, m["refresh"])

}

// Sends request to /refresh, receives new access/refresh tokens pair
func TestRefresh(t *testing.T) {
	var (
		m      map[string]any
		user   map[string]any
		err    error
		cookie map[string]string
	)
	user = make(map[string]any)
	cookie = make(map[string]string)

	user["guid"] = userGuidValid
	body, err := request(signinPath, http.MethodPost, user)
	require.NoError(t, err)
	err = json.Unmarshal(body, &m)
	require.NoError(t, err)
	cookie["refresh"] = fmt.Sprintf("%v", m["refresh"])

	body, err = request(refreshPath, http.MethodPost, nil, cookie)
	require.NoError(t, err)

	err = json.Unmarshal(body, &m)
	require.NoError(t, err)
	assert.NotEmpty(t, m["access"])
	assert.NotEmpty(t, m["refresh"])

}

// Sends request to /refresh with false credentials, receives error
func TestRefreshFalseCredentials(t *testing.T) {
	var (
		m      map[string]any
		user   map[string]any
		err    error
		cookie map[string]string
	)
	user = make(map[string]any)
	cookie = make(map[string]string)

	user["guid"] = userGuidValid
	body, err := request(signinPath, http.MethodPost, user)
	require.NoError(t, err)
	err = json.Unmarshal(body, &m)
	require.NoError(t, err)

	cookie["refresh"] = fmt.Sprintf("%v", m["refresh"])

	body, err = request(refreshPath, http.MethodPost, nil, cookie)
	assert.Error(t, err)

	err = json.Unmarshal(body, &m)
	require.NoError(t, err)
	assert.Empty(t, m["access"])
	assert.Empty(t, m["refresh"])
}
