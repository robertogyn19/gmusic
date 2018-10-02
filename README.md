# Google Music Unofficial API

## 1. Commands

### 1.1. Playlist

```bash
manage playlists

Usage:
  gmusic playlists [command]

Available Commands:
  list        list playlists

Flags:
  -h, --help   help for playlists

Global Flags:
  -e, --email string      email
  -p, --password string   password

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