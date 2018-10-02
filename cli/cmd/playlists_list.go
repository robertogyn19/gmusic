package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/robertogyn19/gmusic"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	email := viper.GetString("email")
	pass := viper.GetString("password")

	conn, err := gmusic.Login(email, pass)
	if err != nil {
		log.Printf("could not login with email %s, error: %v", email, err)
		os.Exit(1)
	}

	playlists, err := conn.ListPlaylists()
	if err != nil {
		log.Printf("could not get list of playlists, error: %v", err)
		os.Exit(1)
	}

	table := playlistsToTable(playlists, []string{"name", "id"})
	fmt.Println(table)
}
