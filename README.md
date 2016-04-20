
# Table of Contents
1. Playlist Routes
  * [Create New Playlist](#create-new-playlist)
  * [Add Song To Playlist](#add-song-to-playlist)
2. Song Routes
  * [Like Song](#like-song)
  * [Delete Song](#delete-song)

# Playlist Routes

#### Create New Playlist
```json
// URL
POST /playlists

// Request
{}

// Response
{}
```
---
#### Add Song to Playlist
```json
// URL
POST /playlists/:playlist_id/songs

// Request
{}

// Response
{}
```

## Song Routes
#### Like Song
```json
// URL
POST /songs/:song_id/like

// Request
{
  "title": <string>,
  "artist": <string>,
  "album": <string>
}

// Response
{
  "songs": [
    {
      "ID": <int>,
      "title": <string>,
      "artist": <string>,
      "album": <string>
    }, {...}, ...
  ]
}
```

---
#### Delete Song
```json
// URL
POST /songs/:song_id/like

// Request
{
  "title": <string>,
  "artist": <string>,
  "album": <string>
}

// Response
{
  "songs": [
    {
      "ID": <int>,
      "title": <string>,
      "artist": <string>,
      "album": <string>
    }, {...}, ...
  ]
}
```
