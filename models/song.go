package models

import (
  "strconv"
)

type Song struct {
  Model
  Title      string       `json:"title"`
  Artist     string       `json:"artist"`
  Album      string       `json:"album"`
  AlbumArt   string       `json:"album_art"`
  SpotifyID  string       `json:"spotify_id"`
  Votes      int          `json:"votes"`
  PlaylistID uint         `json:"playlist_id"`
  ActivePlaylistID uint   `json:"active_playlist_id"`
}

func (song Song) GetData() interface{} {
  return song
}

func SongFromID(id int) Song {
  var song Song
  DB.First(&song, id)

  // Check that song exists
  if DB.NewRecord(song){
    panic(HttpError{404, "Cannot find Song with ID=" + strconv.Itoa(id)})
  }

  return song
}
