package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/robertogyn19/gmusic"
	"github.com/spf13/viper"
)

func getConn() *gmusic.GMusic {
	email := viper.GetString("email")
	pass := viper.GetString("password")

	conn, err := gmusic.Login(email, pass)
	if err != nil {
		log.Printf("could not login with email %s, error: %v", email, err)
		os.Exit(1)
	}

	return conn
}

func playlistsToTable(playlists []*gmusic.Playlist, columns []string) *bytes.Buffer {
	out := bytes.NewBufferString("")
	table := tablewriter.NewWriter(out)

	header := append([]string{"#"}, columns...)
	table.SetHeader(header)

	for i, list := range playlists {
		obj := toMap(list)
		row := getRow(obj, columns)
		idx := valueToString(i + 1)
		row = append([]string{idx}, row...)
		table.Append(row)
	}

	table.Render()
	return out
}

func getRow(obj map[string]interface{}, columns []string) []string {
	row := make([]string, 0)

	for _, c := range columns {
		value := obj[c]
		str := valueToString(value)
		row = append(row, str)
	}

	return row
}

func valueToString(value interface{}) string {
	return fmt.Sprintf("%v", value)
}

func toMap(item interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	d, _ := json.Marshal(item)
	json.Unmarshal(d, &m)
	return m
}
