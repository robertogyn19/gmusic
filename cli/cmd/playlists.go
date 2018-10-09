package cmd

import "github.com/spf13/cobra"

var playlistsCmd = &cobra.Command{
	Use:   "playlists",
	Short: "manage playlists",
}

func init() {
	flags := []Flag{
		{
			Name:  "name",
			Desc:  "playlist name",
			Value: "",
			Short: "n",
		},
	}

	BindFlags(playlistsCmd, flags, true)
	RootCmd.AddCommand(playlistsCmd)
}
