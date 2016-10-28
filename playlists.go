package gmusic

import (
	"encoding/json"
	"io/ioutil"
	"fmt"
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
	//CreationTimestamp     int    `json:"creationTimestamp"`
	//Deleted               bool   `json:"deleted"`
	//LastModifiedTimestamp int    `json:"lastModifiedTimestamp"`
	Name                  string `json:"name"`
	Description           string `json:"description,omitempty"`
	//Type                  string `json:"type"`
	Public                bool   `json:"-"`
	//ShareState            string `json:"-"`
}

type CreatePlaylistResponse struct {
	ID                string        `json:"id"`
	ShareToken        string        `json:"shareToken"`
	CreationTimestamp int64         `json:"creationTimestamp"`
	Track             []interface{} `json:"track"`
}

/*
Params example:
{

}

Response example:
{
	"id": "10bff02f-e438-4bfa-96df-2462da8d5363",
	"sharedToken": "AMaBXynRN-7Co5sCud2a5ff30UKCKtZHBiEKWFna9QpYHH1LYl==",
	"track": [],
	"creationTimestamp":1477579862830000
}
*/
func (g *GMusic) CreatePlaylist(cparams CreatePlaylistParams) (CreatePlaylistResponse, error) {
	ss := "PRIVATE"

	if cparams.Public {
		ss = "PUBLIC"
	}

	params := map[string]interface{}{
		"name": cparams.Name,
		"description": cparams.Description,
		"sharedState": ss,
		"creationTimestamp": 0,
		"deleted": false,
		"lastModifiedTimestamp": -1,
		"type": "USER_GENERATED",
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

	d, err := ioutil.ReadAll(r.Body)

	fmt.Println()
	fmt.Println("--> body", string(d))
	fmt.Println()

	//
	//defer r.Body.Close()
	//if err := json.NewDecoder(r.Body).Decode(&playlist); err != nil {
	//	return nil, err
	//}

	return playlist, nil
}
