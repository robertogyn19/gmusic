package gmusic

import (
	"encoding/json"
)

type ListPlaylists struct {
	Data struct {
		Items []*Playlist `json:"items"`
	} `json:"data"`
	Kind string `json:"kind"`
}

type Playlist struct {
	AccessControlled      bool   `json:"accessControlled"`
	CreationTimestamp     string `json:"creationTimestamp"`
	Deleted               bool   `json:"deleted"`
	ID                    string `json:"id"`
	Kind                  string `json:"kind"`
	LastModifiedTimestamp string `json:"lastModifiedTimestamp"`
	Name                  string `json:"name"`
	OwnerName             string `json:"ownerName"`
	OwnerProfilePhotoUrl  string `json:"ownerProfilePhotoUrl"`
	RecentTimestamp       string `json:"recentTimestamp"`
	ShareToken            string `json:"shareToken"`
	Type                  string `json:"type"`
}

type ListPlaylistEntries struct {
	Data struct {
		Items []*PlaylistEntry `json:"items"`
	} `json:"data"`
	Kind          string `json:"kind"`
	NextPageToken string `json:"nextPageToken"`
}

type PlaylistEntry struct {
	AbsolutePosition      string `json:"absolutePosition"`
	ClientId              string `json:"clientId"`
	CreationTimestamp     string `json:"creationTimestamp"`
	Deleted               bool   `json:"deleted"`
	ID                    string `json:"id"`
	Kind                  string `json:"kind"`
	LastModifiedTimestamp string `json:"lastModifiedTimestamp"`
	PlaylistId            string `json:"playlistId"`
	Source                string `json:"source"`
	TrackId               string `json:"trackId"`
}

func (g *GMusic) ListPlaylists() ([]*Playlist, error) {
	r, err := g.sjRequest("POST", "playlistfeed", nil)
	if err != nil {
		return nil, err
	}
	var data ListPlaylists
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data.Data.Items, nil
}

func (g *GMusic) ListPlaylistEntries() ([]*PlaylistEntry, error) {
	r, err := g.sjRequest("POST", "plentryfeed", nil)
	if err != nil {
		return nil, err
	}
	var data ListPlaylistEntries
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data.Data.Items, nil
}

type CreateMutations struct {
	Mutations []map[string]interface{} `json:"mutations"`
}

type CreatePlaylistParams struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Public      bool   `json:"-"`
}

type CreatePlaylistResponse struct {
	ID           string `json:"id"`
	ClientID     string `json:"client_id"`
	ResponseCode string `json:"response_code"`
}

/*
Params example:
{
	"name": "playlist name",
	"description": "playlist description",
	"sharedState": "PRIVATE | PUBLIC",
	"creationTimestamp": 0,
	"deleted": false,
	"lastModifiedTimestamp": -1,
	"type": "USER_GENERATED",
}

Response example:
{
  "mutate_response": [
    {
      "id": "24e6e72e-0565-40a6-8523-12e1c9090241",
      "client_id": "",
      "response_code": "OK"
    }
  ]
}
*/
func (g *GMusic) CreatePlaylist(cparams CreatePlaylistParams) (CreatePlaylistResponse, error) {
	ss := "PRIVATE"

	if cparams.Public {
		ss = "PUBLIC"
	}

	params := map[string]interface{}{
		"name":                  cparams.Name,
		"description":           cparams.Description,
		"sharedState":           ss,
		"creationTimestamp":     0,
		"deleted":               false,
		"lastModifiedTimestamp": -1,
		"type":                  "USER_GENERATED",
	}

	entry := map[string]interface{}{
		"create": params,
	}

	array := make([]map[string]interface{}, 0)
	array = append(array, entry)

	mutations := CreateMutations{Mutations: array}
	var playlist CreatePlaylistResponse

	r, err := g.sjRequest("POST", "playlistbatch", mutations)

	if err != nil {
		return playlist, err
	}

	defer r.Body.Close()

	var data map[string][]CreatePlaylistResponse

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return playlist, err
	}

	return data["mutate_response"][0], nil
}
