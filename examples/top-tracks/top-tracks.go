package main

import (
	"log"

	"github.com/robertogyn19/gmusic"
	. "github.com/robertogyn19/gmusic/examples"
)

func main() {
	gm := Login()

	pid, err := GetOrCreatePlaylist(gm, "Top Led Musics", "top musics (github.com/robertogyn19/gmusic)")

	if err != nil {
		log.Fatalf("could not get or create playlist: %v", err)
	}

	terms := []string{
		"Led Zeppelin",
	}

	for _, term := range terms {
		nid, err := getArtistId(gm, term)
		if err != nil {
			log.Printf("could not get artist id to %s, error: %v", term, err)
			continue
		}

		artist, err := getArtist(gm, nid)

		if err != nil {
			log.Printf("could not get artist %s, error: %v", term, err)
			continue
		}

		log.Printf("Top tracks of: %s", artist.Name)
		ids := make([]string, len(artist.TopTracks))
		for i, track := range artist.TopTracks {
			log.Printf("%d -- %s", track.TrackNumber, track.Title)
			ids[i] = track.Nid
		}
		log.Println()

		_, err = gm.AddSongsToPlaylist(pid, ids)

		if err != nil {
			log.Printf("could not add tracks to playlist: %v", err)
			continue
		}
	}
}

func GetOrCreatePlaylist(gm *gmusic.GMusic, name, desc string) (string, error) {
	playlists, err := gm.ListPlaylists()

	if err != nil {
		return "", err
	}

	for _, p := range playlists {
		if p.Name == name {
			return p.ID, nil
		}
	}

	params := gmusic.CreatePlaylistParams{Name: name, Description: desc}
	crpr, err := gm.CreatePlaylist(params)

	if err != nil {
		return "", err
	}

	return crpr.ID, nil
}

func getArtistId(gm *gmusic.GMusic, term string) (string, error) {
	opts := gmusic.SearchParams{Term: term, MaxResults: 1}

	sr, err := gm.Search(opts)

	if err != nil {
		return "", err
	}

	for _, entry := range sr.Entries {
		if entry.Type == gmusic.ARTIST_TYPE {
			return entry.Artist.ID, nil
		}
	}

	return "", nil
}

func getArtist(gm *gmusic.GMusic, nid string) (gmusic.Artist, error) {
	params := gmusic.ArtistInfoParams{ID: nid, MaxTopTracts: 15, MaxRelArtist: 1, IncludeAlbums: true}
	return gm.GetArtistInfo(params)
}
