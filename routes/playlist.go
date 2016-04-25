package routes

import (
  "encoding/json"
  "net/http"
  "strconv"
  "fmt"

  "github.com/gorilla/mux"

  "github.com/eecs394-s16/brown-api-golang/models"

)

func addPlaylistRoutes(r *mux.Router) {

  // Get playlist by id
  r.HandleFunc("/playlists/{playlist_id}", getPlaylistHandler).Methods("GET")

  // Create new playlist
  r.HandleFunc("/playlists", createPlaylistHandler).Methods("POST")

  // Add song to playlist
  r.HandleFunc("/playlists/{playlist_id}/songs", addSongToPlaylistHandler).Methods("POST")

  // Delete playlist
  r.HandleFunc("/playlists/{playlist_id}", deletePlaylistHandler).Methods("DELETE")

  // Pop playlist
  r.HandleFunc("/playlists/{playlist_id}/pop", popPlaylistHandler).Methods("PUT")

  // Update playlist
  // TODO
}

func getPlaylistHandler(w http.ResponseWriter, req *http.Request) {
  playlist_id, _ := strconv.Atoi(mux.Vars(req)["playlist_id"])

  // Get playlists
  playlist := models.PlaylistFromID(playlist_id)

  setData(req, playlist.GetData())
}

func createPlaylistHandler(w http.ResponseWriter, req *http.Request) {
  // Get JSON request data
  decoder := json.NewDecoder(req.Body)
  var playlist models.Playlist
  err := decoder.Decode(&playlist)
  if err != nil {
    panic(HttpError{400, err.Error()})
    return
  }

  // Save playlist
  models.DB.Create(&playlist)

  // Return playlist in response
  setData(req, playlist.GetData())
}

func addSongToPlaylistHandler(w http.ResponseWriter, req *http.Request) {
  playlist_id, _ := strconv.Atoi(mux.Vars(req)["playlist_id"])

  // Get JSON request data
  decoder := json.NewDecoder(req.Body)

  // Get playlist from ID
  playlist := models.PlaylistFromID(playlist_id)

  // Create song
  var song models.Song
  err := decoder.Decode(&song)
  if err != nil {
    panic(err)
    return
  }
  // Initialize votes value
  song.Votes = 1
  // Save song
  models.DB.Create(&song)

  fmt.Println(song)

  // Add song to playlist
  models.DB.Model(&playlist).Association("Songs").Append(song)

  // Save playlist
  models.DB.Save(&playlist)

  playlist_update_chan <- playlist_id
  setData(req, playlist.GetData())
}

func deletePlaylistHandler(w http.ResponseWriter, req *http.Request) {
  playlist_id, _ := strconv.Atoi(mux.Vars(req)["playlist_id"])

  // Get playlist from ID
  playlist := models.PlaylistFromID(playlist_id)

  // Delete playlist
  playlist.Delete()

  data := make(map[string]interface{})
  data["deleted"] = playlist_id

  playlist_update_chan <- playlist_id
  setData(req, data)
}

func popPlaylistHandler(w http.ResponseWriter, req *http.Request) {
  playlist_id, _ := strconv.Atoi(mux.Vars(req)["playlist_id"])

  // Get playlist from ID
  playlist := models.PlaylistFromID(playlist_id)

  // Get old and new active song
  var old_active_song models.Song
  var new_active_song models.Song
  models.DB.Model(&playlist).Association("ActiveSong").Find(&old_active_song)
  models.DB.Order("Votes desc").Limit(1).Model(&playlist).Association("Songs").Find(&new_active_song)

  // Put old active song back in queue with zero votes
  if !models.DB.NewRecord(old_active_song) {
    old_active_song.Votes = 0
    models.DB.Save(&old_active_song)

    models.DB.Model(&playlist).Association("Songs").Append(old_active_song)
    models.DB.Model(&playlist).Association("ActiveSong").Delete(old_active_song)
  }

  // Set new active song
  models.DB.Model(&playlist).Association("ActiveSong").Append(new_active_song)
  models.DB.Model(&playlist).Association("Songs").Delete(new_active_song)

  playlist_update_chan <- playlist_id
  setData(req, playlist.GetData())
}
