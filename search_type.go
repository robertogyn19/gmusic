package gmusic


type SearchType string

// 1: Song, 2: Artist, 3: Album, 4: Playlist, 5: Genre,
// 6: Station, 7: Situation, 8: Video, 9: Podcast Series
const (
	SongType      SearchType = "1"
	ArtistType    SearchType = "2"
	AlbumType     SearchType = "3"
	PlaylistType  SearchType = "4"
	GenreType     SearchType = "5"
	StationType   SearchType = "6"
	SituationType SearchType = "7"
	VideoType     SearchType = "8"
	PodcastType   SearchType = "9"
)

func (st SearchType) String() string {
	return string(st)
}
