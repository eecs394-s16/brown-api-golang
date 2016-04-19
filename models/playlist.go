package models

import (
  "strconv"
)

type Playlist struct {
  Model
  Name  string  `json:"name"`
  Songs []Song  `json:"songs"`
}

func (playlist Playlist) GetData() interface{} {
  var songs []Song
  DB.Model(&playlist).Association("Songs").Find(&songs)

  data := make(map[string]interface{})
  data["name"]  = playlist.Name
  data["songs"] = songs
  data["id"] = playlist.ID

  return data
}

func PlaylistFromID(id int) Playlist {
  var playlist Playlist
  DB.First(&playlist, id)

  // Check that playlist exists
  if DB.NewRecord(playlist){
    panic(HttpError{404, "Cannot find Playlist with ID=" + strconv.Itoa(id)})
  }

  return playlist
}
