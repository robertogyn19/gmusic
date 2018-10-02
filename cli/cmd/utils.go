package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/olekukonko/tablewriter"
	"github.com/robertogyn19/gmusic"
)

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
