package gmusic

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type ListTracks struct {
	Data struct {
		Items []*Track `json:"items"`
	} `json:"data"`
	Kind          string `json:"kind"`
	NextPageToken string `json:"nextPageToken"`
}

type Track struct {
	Album       string `json:"album"`
	AlbumArtRef []struct {
		URL string `json:"url"`
	} `json:"albumArtRef"`
	AlbumArtist  string `json:"albumArtist"`
	AlbumId      string `json:"albumId"`
	Artist       string `json:"artist"`
	ArtistArtRef []struct {
		URL string `json:"url"`
	} `json:"artistArtRef"`
	ArtistId              []string `json:"artistId"`
	ClientId              string   `json:"clientId"`
	CreationTimestamp     string   `json:"creationTimestamp"`
	Deleted               bool     `json:"deleted"`
	DiscNumber            int      `json:"discNumber"`
	DurationMillis        string   `json:"durationMillis"`
	EstimatedSize         string   `json:"estimatedSize"`
	ID                    string   `json:"id"`
	Kind                  string   `json:"kind"`
	LastModifiedTimestamp string   `json:"lastModifiedTimestamp"`
	Nid                   string   `json:"nid"`
	PlayCount             int      `json:"playCount"`
	RecentTimestamp       string   `json:"recentTimestamp"`
	StoreId               string   `json:"storeId"`
	Title                 string   `json:"title"`
	TrackNumber           int      `json:"trackNumber"`
	TrackType             string   `json:"trackType"`
	Year                  int      `json:"year"`
}

func (g *GMusic) ListTracks() ([]*Track, error) {
	var tracks []*Track
	var next string
	for {
		r, err := g.sjRequest(http.MethodPost, "trackfeed", struct {
			StartToken string `json:"start-token"`
		}{
			StartToken: next,
		})
		if err != nil {
			return nil, err
		}
		var data ListTracks
		defer r.Body.Close()

		body, _ := ioutil.ReadAll(r.Body)
		if err := json.Unmarshal(body, &data); err != nil {
			return nil, err
		}

		tracks = append(tracks, data.Data.Items...)
		next = data.NextPageToken
		if next == "" || len(tracks) >= 1000 {
			break
		}
	}
	return tracks, nil
}
