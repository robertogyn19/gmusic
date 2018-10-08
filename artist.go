package gmusic

import (
	"encoding/json"
	"net/http"
)

type ArtistInfoParams struct {
	ID            string `url:"nid"`
	IncludeAlbums bool   `url:"include-albums"`
	MaxTopTracts  int    `url:"num-top-tracks"`
	MaxRelArtist  int    `url:"num-related-artists"`
	Alt           string `url:"alt"`
}

func (g *GMusic) GetArtistInfo(params ArtistInfoParams) (Artist, error) {
	art := Artist{}
	params.Alt = "json"

	r, err := g.sjRequest(http.MethodGet, "fetchartist", params)

	if err != nil {
		return art, err
	}

	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&art); err != nil {
		return art, err
	}

	return art, nil
}
