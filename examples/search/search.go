package main

import (
	"log"

	"github.com/robertogyn19/gmusic"
	. "github.com/robertogyn19/gmusic/examples"
)

func main() {
	gm := Login()

	sr, err := gm.Search(gmusic.SearchParams{Term: "WINNER"})

	if err != nil {
		log.Fatal("Search error ", err)
	}

	for _, entry := range sr.Entries {
		switch entry.Type {
		case "1": // Track
			t := entry.Track
			log.Printf("(%s) Track info --> Artist: %s | Album: %s | TrackNumber: %d | Year: %d", entry.Type, t.Artist, t.Album, t.TrackNumber, t.Year)
		case "2": // Artist
			a := entry.Artist
			log.Printf("(%s) Artist info --> Artist name: %s", entry.Type, a.Name)
		case "3": // Album
			a := entry.Album
			log.Printf("(%s) Album info --> Album name: %s | Artist name: %s", entry.Type, a.Name, a.Artist)
		case "4": // Playlist
			p := entry.Playlist
			log.Printf("(%s) Playlist info --> Name: %s | Owner: %s", entry.Type, p.Name, p.OwnerName)
		}
		log.Println()
	}
}
