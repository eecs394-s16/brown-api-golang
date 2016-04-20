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
  setData(req, data)
}
