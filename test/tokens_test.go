package test

import "testing"

// Sends request to /signin, recieves access/refresh tokens if user is correct
func TestSignin(t *testing.T) {
}

// Sends request to /signin with false credentials, receives error
func TestSigninFalseCredentials(t *testing.T) {

}

// Sends request to /refresh, receives new access/refresh tokens pair
func TestRefresh(t *testing.T) {

}

// Sends request to /refresh with false credentials, receives error
func TestRefreshFalseCredentials(t *testing.T) {
}
