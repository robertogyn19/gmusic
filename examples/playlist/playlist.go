package main

import (
	"log"

	"github.com/robertogyn19/gmusic"
	. "github.com/robertogyn19/gmusic/examples"
)

func main() {
	gm := Login()

	cpp := gmusic.CreatePlaylistParams{Name: "new playlist with gmusic"}
	pl, err := gm.CreatePlaylist(cpp)

	if err != nil {
		log.Fatal("Creating playlist error ", err)
	}

	log.Println("Playlist id:", pl.ID)
}

