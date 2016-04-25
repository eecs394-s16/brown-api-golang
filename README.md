These routes are **not live** yet! Keep using previous routes.

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
{
  "name": <string>
}

// Response
{
  "id": <int>,
  "name": <string>,
  "songs": [{
      "id": <int>,
      "title": <string>,
      "artist": <string>,
      "album": <string>,
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
  "spotify_id": <string>
}

// Response
{
  "id": <int>,
  "title": <string>,
  "artist": <string>,
  "album": <string>,
  "spotify_id": <string>
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
