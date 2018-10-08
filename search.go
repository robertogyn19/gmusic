package gmusic

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/go-querystring/query"
)

type SearchType string

// 1: Song, 2: Artist, 3: Album, 4: Playlist, 5: Genre,
// 6: Station, 7: Situation, 8: Video, 9: Podcast Series
const (
	SongType      SearchType = "1"
	ArtistType    SearchType = "2"
	AlbumType     SearchType = "3"
	PlaylistType  SearchType = "4"
	GenreType     SearchType = "5"
	StationType   SearchType = "6"
	SituationType SearchType = "7"
	VideoType     SearchType = "8"
	PodcastType   SearchType = "9"
)

type SearchParams struct {
	Term       string `url:"q"`
	MaxResults int    `url:"max-results"`
}

type SearchResponse struct {
	Kind          string          `json:"kind"`
	ClusterDetail []ClusterDetail `json:"clusterDetail"`
}

type ClusterDetail struct {
	ResultToken string                  `json:"resultToken"`
	Cluster     ClusterResult           `json:"cluster"`
	Entries     []SearchEntriesResponse `json:"entries"`
}

type ClusterResult struct {
	Type     SearchType `json:"type"`
	Category string     `json:"category"`
	Id       string     `json:"search_genre"`
}

type SearchEntriesResponse struct {
	Type               SearchType `json:"type"`
	Artist             Artist     `json:"artist,omitempty"`
	Album              Album      `json:"album,omitempty"`
	Track              Track      `json:"track,omitempty"`
	Playlist           Playlist   `json:"playlist,omitempty"`
	BestResult         bool       `json:"best_result"`
	NavigationalResult bool       `json:"navigational_result"`
}

/*
JSON example:

{
  "kind": "sj#artist",
  "name": "WINNER",
  "artistArtRef": "http://lh3.googleusercontent.com/sU8NsTudlWZ4TQP2hgNQDjkN3RM0xGy5J9k8m3G6LeAX0Yk4hoXrzLZEkgfkVTWvGX9taawz0ao",
  "artistArtRefs": [
    {
      "kind": "sj#imageRef",
      "url": "http://lh3.googleusercontent.com/sU8NsTudlWZ4TQP2hgNQDjkN3RM0xGy5J9k8m3G6LeAX0Yk4hoXrzLZEkgfkVTWvGX9taawz0ao",
      "aspectRatio": "2",
      "autogen": false
    }
  ],
  "artistId": "Aglg43ajc3toter3svcvwjp3vky",
  "artist_bio_attribution": {
    "kind": "sj#attribution",
    "source_title": "artist representative"
  }
}
*/
type Artist struct {
	ID                   string               `json:"artistId"`
	Name                 string               `json:"name"`
	ArtistArtRef         string               `json:"artistArtRef"`
	ArtistArtRefs        []ArtRefs            `json:"artRefs"`
	ArtistBioAttribution ArtistBioAttribution `json:"artist_bio_attribution"`
	TopTracks            []Track              `json:"topTracks"`
}

type ArtistBioAttribution struct {
	Kind        string `json:"kind"`
	SourceTitle string `json:"source_title"`
}

type ArtRefs struct {
	Kind        string `json:"kind"`
	URL         string `json:"url"`
	AspectRatio string `json:"aspectRatio"`
	AutoGen     bool   `json:"autogen"`
}

/*
JSON example:

{
  "kind": "sj#album",
  "name": "2014 S/S",
  "albumArtist": "WINNER",
  "albumArtRef": "http://lh3.googleusercontent.com/oIRCWf0HS-RYw2jhU4deDqgAoWAyJJHGkpEZUi8qqz09aTXuww1W7Qe7AlT56mofoctXasEguQ",
  "albumId": "Bsdo22syyl2p2s5jqasgexmgph4",
  "artist": "WINNER",
  "artistId": [
    "Aglg43ajc3toter3svcvwjp3vky"
  ],
  "description_attribution": {
    "kind": "sj#attribution",
    "source_title": "Wikipedia",
    "source_url": "https://en.wikipedia.org/wiki/2014_S/S",
    "license_title": "Creative Commons Attribution CC-BY-SA 4.0",
    "license_url": "http://creativecommons.org/licenses/by-sa/4.0/legalcode"
  },
  "year": 2014,
  "explicitType": "2"
}
*/
type Album struct {
	Kind                   string                 `json:"kind"`
	Name                   string                 `json:"name"`
	AlbumArtist            string                 `json:"albumArtist"`
	AlbumArtRef            string                 `json:"albumArtRef"`
	AlbumID                string                 `json:"albumId"`
	Artist                 string                 `json:"artist"`
	ArtistID               []string               `json:"artistId"`
	DescriptionAttribution DescriptionAttribution `json:"description_attribution"`
	Year                   int                    `json:"year"`
	ExplicitType           string                 `json:"explicitType"`
}

type DescriptionAttribution struct {
	Kind         string `json:"kind"`
	SourceTitle  string `json:"source_title"`
	SourceURL    string `json:"source_url"`
	LicenseTitle string `json:"license_title"`
	LicenseURL   string `json:"license_url"`
}

const (
	ctParam = "1,2,3,4,5,6,7,8,9"
)

func (g *GMusic) Search(opts SearchParams) (SearchResponse, error) {
	var sr SearchResponse

	if opts.MaxResults == 0 {
		opts.MaxResults = 1
	}

	params, err := query.Values(opts)

	if err != nil {
		return sr, err
	}

	url := fmt.Sprintf("query?%s", params.Encode())
	url = fmt.Sprintf("%s&ct=%s&ic=true", url, ctParam)
	r, err := g.sjRequest(http.MethodGet, url, nil)

	if err != nil {
		return sr, err
	}

	reader := r.Body
	body, _ := ioutil.ReadAll(reader)
	defer reader.Close()

	if err := json.Unmarshal(body, &sr); err != nil {
		return sr, err
	}

	return sr, nil
}
