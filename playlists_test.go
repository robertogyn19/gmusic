package gmusic

import (
	"testing"
	"net/http"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestGMusic_CreatePlaylist(t *testing.T) {
	defer gock.Off()
	gock.DisableNetworking()

	gtest := new(GMusic)

	params := CreatePlaylistParams{
		Name: "test 1",
		Public: false,
		Description: "create playlist test 1",
	}

	reqMock := gock.New(sjURL)
	reqMock.Post("playlistbatch").Reply(http.StatusOK).BodyString(`{}`)

	pmr, err := gtest.CreatePlaylist(params)
	assert.NoError(t, err)
	assert.NotEmpty(t, pmr.ResponseCode)
}
