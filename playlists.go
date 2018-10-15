package gmusic

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/satori/go.uuid"
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
	r, err := g.sjRequest(http.MethodPost, "playlistfeed", nil)
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
	r, err := g.sjRequest(http.MethodPost, "plentryfeed", nil)
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

type PlaylistMutateResponse struct {
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
func (g *GMusic) CreatePlaylist(cparams CreatePlaylistParams) (PlaylistMutateResponse, error) {
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
	var playlist PlaylistMutateResponse

	r, err := g.sjRequest(http.MethodPost, "playlistbatch", mutations)

	if err != nil {
		return playlist, err
	}

	defer r.Body.Close()

	var data map[string][]PlaylistMutateResponse

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return playlist, err
	}

	return data["mutate_response"][0], nil
}

/**
Response example:
{
  "mutate_response": [
    {
      "id": "dd2197a2-3f19-3668-874f-512760f84a62",
      "client_id": "39a55b1c-403a-4cc1-a0a0-aace50ca2d2e",
      "response_code": "OK"
    },
    {
      "id": "b30bebd8-fc23-36e8-8c41-45f127f49baa",
      "client_id": "26690b16-b6fb-41fa-becb-cfa6ef90f358",
      "response_code": "OK"
    },
    {
      "id": "4e7154c9-3eda-3bf1-b62b-d98c43094c7e",
      "client_id": "dffb83f1-a42e-4ed8-8180-1197c75fc7df",
      "response_code": "OK"
    },
    {
      "id": "2f6cbcd3-a233-31e9-9989-1f4f26e69515",
      "client_id": "d3d81efc-1561-429e-aab6-544ce8ae1e92",
      "response_code": "OK"
    },
    {
      "id": "6dd36f02-681f-37a3-bad3-f18d7c9aae8f",
      "client_id": "c2913f61-ca98-4599-863a-0de9f7b53fdb",
      "response_code": "OK"
    }
  ]
}
*/
func (g *GMusic) AddSongsToPlaylist(pid string, trackIds []string) (PlaylistMutateResponse, error) {
	entries := make([]map[string]interface{}, 0)

	prev, curr, next := "", uuid.NewV4().String(), uuid.NewV4().String()

	for i, trackId := range trackIds {
		m := map[string]interface{}{
			"clientId":              curr,
			"creationTimestamp":     "-1",
			"deleted":               false,
			"lastModifiedTimestamp": "0",
			"playlistId":            pid,
			"source":                1,
			"trackId":               trackId,
		}

		if strings.HasPrefix(trackId, "T") {
			m["source"] = 2
		}

		if i > 0 {
			m["precedingEntryId"] = prev
		}

		if i < (len(trackIds) - 1) {
			m["followingEntryId"] = next
		}

		entry := map[string]interface{}{
			"create": m,
		}

		entries = append(entries, entry)
		prev, curr, next = curr, next, uuid.NewV4().String()
	}

	mutations := CreateMutations{Mutations: entries}

	r, err := g.sjRequest(http.MethodPost, "plentriesbatch?alt=json", mutations)

	if err != nil {
		return PlaylistMutateResponse{}, err
	}

	defer r.Body.Close()

	var data map[string][]PlaylistMutateResponse

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return PlaylistMutateResponse{}, err
	}

	return data["mutate_response"][0], nil
}
