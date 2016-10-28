package main

import (
	"log"
	"os"

	"github.com/robertogyn19/gmusic"
)

func main() {
	user := os.Getenv("GOOGLE_USER")
	pass := os.Getenv("GOOGLE_PASS")

	if user == "" || pass == "" {
		log.Fatal("Invalid credentials")
	}

	gm, err := gmusic.Login(user, pass)

	if err != nil {
		log.Fatal("Login error ", err)
	}

	cpp := gmusic.CreatePlaylistParams{Name: "new playlist with gmusic"}
	gm.CreatePlaylist(cpp)
	log.Printf("Logged as %s", gm.DeviceID)
}
