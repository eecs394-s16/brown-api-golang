# Table of Contents
1. Playlist Routes
  * [Create New Playlist](#create-new-playlist)
  * [Add Song To Playlist](#add-song-to-playlist)
  * [Pop Playlist](#pop-playlist)
2. Song Routes
  * [Like Song](#like-song)
  * [Delete Song](#delete-song)

# Playlist Routes

#### Create New Playlist
```json
// URL
POST /playlists

// Request
{
  "name": <string>
}

// Response
{
  "id": <int>,
  "name": <string>,
  "active_song": {
    "id": <int>,
    "title": <string>,
    "artist": <string>,
    "album": <string>,
    "album_art": <string>,
    "spotify_id": <string>
  }
  "songs": [{
    "id": <int>,
    "title": <string>,
    "artist": <string>,
    "album": <string>,
    "album_art": <string>,
    "spotify_id": <string>
  }, ...]
}
```

#### Add Song to Playlist
```json
// URL
POST /playlists/:playlist_id/songs

// Request
{
  "title": <string>,
  "artist": <string>,
  "album": <string>,
  "album_art": <string>,
  "spotify_id": <string>
}

// Response
{
  "id": <int>,
  "title": <string>,
  "artist": <string>,
  "album": <string>,
  "album_art": <string>,
  "spotify_id": <string>
}
```

#### Pop Playlist
```json
// URL
PUT /playlists/:playlist_id/pop

// Request
{}

// Response
{
  "id": <int>,
  "name": <string>,
  "active_song": {
    "id": <int>,
    "title": <string>,
    "artist": <string>,
    "album": <string>,
    "album_art": <string>,
    "spotify_id": <string>
  }
  "songs": [{
    "id": <int>,
    "title": <string>,
    "artist": <string>,
    "album": <string>,
    "album_art": <string>,
    "spotify_id": <string>
  }, ...]
}
```

## Song Routes
#### Like Song
```json
// URL
POST /songs/:song_id/like

// Request
{}

// Response
{}

```

#### Delete Song
```json
// URL
POST /songs/:song_id/like

// Request
{}

// Response
{}
```
