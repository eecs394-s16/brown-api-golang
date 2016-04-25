package models

import (
  "strconv"
)

type Playlist struct {
  Model
  Name       string  `json:"name"`
  Songs      []Song  `json:"songs"       gorm:"ForeignKey:PlaylistID"`
  ActiveSong Song    `json:"active_song" gorm:"ForeignKey:ActivePlaylistID"`
}

func (playlist Playlist) GetData() interface{} {
  var songs []Song
  var active_song Song
  DB.Order("Votes desc").Model(&playlist).Association("Songs").Find(&songs)
  DB.Model(&playlist).Association("ActiveSong").Find(&active_song)

  data := make(map[string]interface{})
  data["name"]  = playlist.Name
  data["songs"] = songs
  data["active_song"] = active_song
  data["id"]    = playlist.ID

  if DB.NewRecord(active_song){
    data["active_song"] = nil
  }

  return data
}

func (playlist *Playlist) Delete() {
  // Loop through each song in playlist and delete
  var songs []Song
  DB.Model(&playlist).Association("Songs").Find(&songs)
  for _, song := range songs {
    DB.Delete(&song)
  }

  // Delete playlist
  DB.Delete(&playlist)
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
