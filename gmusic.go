// Package gmusic provides methods to list music, playlists, artists, etc
// from Google Play Music.
package gmusic

import (
	"github.com/mjibson/gpsoauth"
)

const (
	googlePlayMusicEndpoint = "https://play.google.com/music"
	serviceName             = "sj"
	sjURL                   = "https://www.googleapis.com/sj/v2.4/"
)

type GMusic struct {
	DeviceID string
	Auth     string
}

// Login logs in with a username and password and androidID from a MAC
// address of the machine.
func Login(username, password string) (*GMusic, error) {
	return LoginAndroid(username, password, gpsoauth.GetNode())
}

// LoginAndroid logs in with a username and password and given androidID.
func LoginAndroid(username, password, androidID string) (*GMusic, error) {
	auth, err := gpsoauth.Login(username, password, androidID, serviceName)
	if err != nil {
		return nil, err
	}
	gm := GMusic{
		Auth: auth,
	}
	if err := gm.setDeviceID(); err != nil {
		return nil, err
	}
	return &gm, nil
}
