package cmd

import "github.com/spf13/cobra"

var playlistsCmd = &cobra.Command{
	Use:   "playlists",
	Short: "manage playlists",
}

func init() {
	RootCmd.AddCommand(playlistsCmd)
}

