# Google Music Unofficial API

## 1. Commands

### 1.1. Playlist

```bash
manage playlists

Usage:
  gmusic playlists [command]

Available Commands:
  create      create playlist with top tracks
  list        list playlists

Flags:
  -h, --help          help for playlists
  -n, --name string   playlist name

Global Flags:
  -e, --email string      email
  -p, --password string   password
  -v, --verbose           verbose

Use "gmusic playlists [command] --help" for more information about a command.
```

#### 1.1.1. List

```bash
go run cli/main.go playlists list -e $GOOGLE_EMAIL -p $GOOGLE_PASS
+----+--------------------------------+--------------------------------------+
| #  |              NAME              |                  ID                  |
+----+--------------------------------+--------------------------------------+
|  1 | Nacional                       | 76982861-ee43-4305-8af4-7686b6443e9a |
|  2 | The One                        | 8586a9d7-7aa8-3d56-9476-b9b627b00054 |
+----+--------------------------------+--------------------------------------+
```

#### 1.1.2. Create

```bash
go run cli/main.go playlists create -e $GOOGLE_EMAIL -p $GOOGLE_PASS -n "Good Rock" -a "Led Zeppelin" -a "ACDC" -a "Guns N' Roses" -v

2018/10/09 07:24:21 artist: Led Zeppelin
2018/10/09 07:24:21 1  - Immigrant Song
2018/10/09 07:24:21 2  - Stairway To Heaven
2018/10/09 07:24:21 3  - Rock And Roll
2018/10/09 07:24:21 4  - Black Dog
2018/10/09 07:24:21 5  - Ramble On
2018/10/09 07:24:21 6  - Whole Lotta Love
2018/10/09 07:24:21 7  - Going To California
2018/10/09 07:24:21 8  - Kashmir
2018/10/09 07:24:21 9  - Good Times Bad Times
2018/10/09 07:24:21 10 - Communication Breakdown
2018/10/09 07:24:21 11 - All My Love
2018/10/09 07:24:21 12 - When The Levee Breaks
2018/10/09 07:24:21 13 - Babe I'm Gonna Leave You
2018/10/09 07:24:21 14 - Heartbreaker
2018/10/09 07:24:21 15 - D'yer Mak'er
2018/10/09 07:24:21 
2018/10/09 07:24:21 artist: AC/DC
2018/10/09 07:24:21 1  - Back In Black
2018/10/09 07:24:21 2  - Thunderstruck
2018/10/09 07:24:21 3  - You Shook Me All Night Long
2018/10/09 07:24:21 4  - Hells Bells
2018/10/09 07:24:21 5  - T.N.T.
2018/10/09 07:24:21 6  - Shoot to Thrill
2018/10/09 07:24:21 7  - Dirty Deeds Done Dirt Cheap
2018/10/09 07:24:21 8  - For Those About to Rock (We Salute You)
2018/10/09 07:24:21 9  - Have a Drink on Me
2018/10/09 07:24:21 10 - Moneytalks
2018/10/09 07:24:21 11 - Rock and Roll Ain't Noise Pollution
2018/10/09 07:24:21 12 - It's a Long Way to the Top (If You Wanna Rock 'N' Roll)
2018/10/09 07:24:21 13 - Who Made Who
2018/10/09 07:24:21 14 - What Do You Do for Money Honey
2018/10/09 07:24:21 15 - If You Want Blood (You've Got It)
2018/10/09 07:24:21 
2018/10/09 07:24:21 artist: Guns N' Roses
2018/10/09 07:24:21 1  - November Rain
2018/10/09 07:24:21 2  - Patience
2018/10/09 07:24:21 3  - Don't Cry
2018/10/09 07:24:21 4  - Sweet Child O' Mine
2018/10/09 07:24:21 5  - Live And Let Die
2018/10/09 07:24:21 6  - Paradise City
2018/10/09 07:24:21 7  - Welcome To The Jungle
2018/10/09 07:24:21 8  - Knockin' On Heaven's Door
2018/10/09 07:24:21 9  - You Could Be Mine
2018/10/09 07:24:21 10 - Civil War
2018/10/09 07:24:21 11 - Yesterdays
2018/10/09 07:24:21 12 - Estranged
2018/10/09 07:24:21 13 - Nightrain
2018/10/09 07:24:21 14 - Back Off Bitch (Album Version Explicit)
2018/10/09 07:24:21 15 - 14 Years
2018/10/09 07:24:21 
2018/10/09 07:24:21 45 songs were successfully added to playlist Good Rock
```