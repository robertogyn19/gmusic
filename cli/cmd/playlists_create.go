package cmd

import (
	"log"
	"os"

	"github.com/robertogyn19/gmusic"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var playlistsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create playlist with top tracks",
	Run:   playlistsCreateCmdRun,
}

func init() {
	flags := []Flag{

		{
			Name:  "artists",
			Desc:  "artist name",
			Value: []string{},
			Short: "a",
		},
	}

	BindFlags(playlistsCreateCmd, flags, true)
	playlistsCmd.AddCommand(playlistsCreateCmd)
}

func playlistsCreateCmdRun(cmd *cobra.Command, _ []string) {
	conn := getConn()

	playlistName := viper.GetString("name")
	terms := viper.GetStringSlice("artists")
	if playlistName == "" || len(terms) == 0 {
		cmd.Usage()
		os.Exit(1)
	}

	cparams := gmusic.CreatePlaylistParams{
		Name:        playlistName,
		Description: "playlist created with github.com/robertogyn19/gmusic",
	}

	pmr, err := conn.CreatePlaylist(cparams)
	if err != nil {
		log.Printf("could not create playlist with name %s, error: %v", err)
		os.Exit(1)
	}

	artists, err := conn.SearchArtists(terms)
	if err != nil {
		os.Exit(1)
	}

	trackIds := make([]string, 0)
	for _, artist := range artists {
		if verbose {
			log.Printf("artist: %s", artist.Name)
		}
		for i, tt := range artist.TopTracks {
			if verbose {
				log.Printf("%-2d - %s", i+1, tt.Title)
			}
			trackIds = append(trackIds, tt.Nid)
		}

		if verbose {
			log.Println()
		}
	}

	pmr, err = conn.AddSongsToPlaylist(pmr.ID, trackIds)
	if err != nil {
		log.Printf("could not add songs to playlist %s, error: %v", playlistName, err)
		os.Exit(1)
	}

	log.Printf("%d songs were successfully added to playlist %s", len(trackIds), playlistName)
}
