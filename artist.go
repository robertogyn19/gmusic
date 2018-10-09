package gmusic

import (
	"encoding/json"
	"log"
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

func (g *GMusic) SearchArtists(terms []string) ([]Artist, error) {
	list := make([]Artist, 0)

	for _, artist := range terms {
		searchParam := SearchParams{
			Term:        artist,
			MaxResults:  1,
			SearchTypes: []SearchType{ArtistType},
		}

		response, err := g.Search(searchParam)
		if err != nil {
			log.Printf("could not search for artist %s, error: %v", artist, err)
			return list, err
		}

		for _, detail := range response.ClusterDetail {
			if detail.Cluster.Type == ArtistType {
				for _, entry := range detail.Entries {
					list = append(list, entry.Artist)
				}
			}
		}
	}

	var err error
	for i, artist := range list {
		params := ArtistInfoParams{
			ID:            artist.ID,
			MaxTopTracts:  15,
			MaxRelArtist:  0,
			IncludeAlbums: false,
		}
		artist, err = g.GetArtistInfo(params)
		if err != nil {
			log.Printf("could not get artist info for %s, error: %v", artist.Name, err)
			continue
		}

		list[i] = artist
	}

	return list, err
}
