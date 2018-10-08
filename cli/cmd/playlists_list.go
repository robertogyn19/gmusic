package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var playlistsListCmd = &cobra.Command{
	Use:   "list",
	Short: "list playlists",
	Run:   playlistsListCmdRun,
}

func init() {
	playlistsCmd.AddCommand(playlistsListCmd)
}

func playlistsListCmdRun(_ *cobra.Command, _ []string) {
	conn := getConn()

	playlists, err := conn.ListPlaylists()
	if err != nil {
		log.Printf("could not get list of playlists, error: %v", err)
		os.Exit(1)
	}

	table := playlistsToTable(playlists, []string{"name", "id"})
	fmt.Println(table)
}
