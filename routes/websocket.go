package routes

import (
  _ "encoding/json"
  "net/http"
  "fmt"
  "strconv"
  "time"

  "github.com/gorilla/mux"
  "github.com/gorilla/websocket"

  "github.com/eecs394-s16/brown-api-golang/models"
)

type newPlaylistConn struct {
  c *websocket.Conn
  playlist_id int
}

var playlist_add_chan    = make(chan newPlaylistConn)
var playlist_rem_chan    = make(chan newPlaylistConn)
var playlist_update_chan = make(chan int)

func addWebsocketRoutes(r *mux.Router) {
  go playlistConnectionManager()

  go func() {
    for {
      time.Sleep(1*time.Second)
      playlist_update_chan <- 1
      fmt.Println("test")
    }
  }()

  r.HandleFunc("/ws/playlists/{playlist_id}", wsPlaylistHandler)
}

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}

func playlistConnectionManager() {
  connection_map := make(map[int][]*websocket.Conn)
  for {
    select {

    // Handle opened connection
    // =========================
    case new_conn := <- playlist_add_chan:
      connection_map[new_conn.playlist_id] = append(connection_map[new_conn.playlist_id], new_conn.c)

    // Handle closed connection
    // =========================
    case rem_conn := <- playlist_rem_chan:
      // Loop through connections to playlist
      for i, conn := range connection_map[rem_conn.playlist_id] {
        // If connection matches passed in connection...
        if conn == rem_conn.c {
          // Remove connection from slice
          connection_map[rem_conn.playlist_id] = append(connection_map[rem_conn.playlist_id][:i], connection_map[rem_conn.playlist_id][i+1:]...)
          break
        }
      }

    // Handle updated playlist
    // ========================
    case playlist_id := <- playlist_update_chan:
      playlist := models.PlaylistFromID(playlist_id)

      for _, c := range connection_map[playlist_id] {
        c.WriteJSON(&playlist)
      }
    }

  }
}

func wsPlaylistHandler(w http.ResponseWriter, req *http.Request) {
  playlist_id, _ := strconv.Atoi(mux.Vars(req)["playlist_id"])
  fmt.Println(playlist_id)

  c, err := upgrader.Upgrade(w, req, nil)
  if err != nil {
		fmt.Println("upgrade:", err)
		return
	}
  defer func() {
    playlist_rem_chan <- newPlaylistConn{c, playlist_id}
    c.Close()
  }()

  // Add new connection to playlistConnectionManager
  playlist_add_chan <- newPlaylistConn{c, playlist_id}

  for {
    _, message, err := c.ReadMessage()
    if err != nil {
      return
    }
    fmt.Println(message)
    c.WriteJSON(message)
  }

}
