package gmusic

import (
	"encoding/json"
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
	DiscNumber            float64  `json:"discNumber"`
	DurationMillis        string   `json:"durationMillis"`
	EstimatedSize         string   `json:"estimatedSize"`
	ID                    string   `json:"id"`
	Kind                  string   `json:"kind"`
	LastModifiedTimestamp string   `json:"lastModifiedTimestamp"`
	Nid                   string   `json:"nid"`
	PlayCount             float64  `json:"playCount"`
	RecentTimestamp       string   `json:"recentTimestamp"`
	StoreId               string   `json:"storeId"`
	Title                 string   `json:"title"`
	TrackNumber           float64  `json:"trackNumber"`
	TrackType             string   `json:"trackType"`
	Year                  float64  `json:"year"`
}

func (g *GMusic) ListTracks() ([]*Track, error) {
	var tracks []*Track
	var next string
	for {
		r, err := g.sjRequest("POST", "trackfeed", struct {
			StartToken string `json:"start-token"`
		}{
			StartToken: next,
		})
		if err != nil {
			return nil, err
		}
		var data ListTracks
		defer r.Body.Close()
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			return nil, err
		}
		tracks = append(tracks, data.Data.Items...)
		next = data.NextPageToken
		if next == "" {
			break
		}
	}
	return tracks, nil
}
