package gmusic

import (
	"io/ioutil"
	"testing"
)

func TestGMusic(t *testing.T) {
	username, err := ioutil.ReadFile("username")
	if err != nil {
		t.Fatal(err)
	}
	password, err := ioutil.ReadFile("password")
	if err != nil {
		t.Fatal(err)
	}
	gm, err := Login(string(username), string(password))
	if err != nil {
		t.Fatal(err)
	}
	_, err = gm.ListPlaylists()
	if err != nil {
		t.Fatal(err)
	}
	_, err = gm.ListPlaylistEntries()
	if err != nil {
		t.Fatal(err)
	}
	_, err = gm.ListTracks()
	if err != nil {
		t.Fatal(err)
	}
}
