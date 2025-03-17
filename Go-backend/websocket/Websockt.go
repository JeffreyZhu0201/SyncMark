package websocket

import (
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}
var rooms = make(map[string][]*websocket.Conn)
var roomsMutex = &sync.Mutex{}

func HandleWebSocket(w http.ResponseWriter, r *http.Request, roomId string) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer conn.Close()

	roomsMutex.Lock()
	rooms[roomId] = append(rooms[roomId], conn)
	roomsMutex.Unlock()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			break
		}
		BroadcastMessage(roomId, message)
	}
}

func BroadcastMessage(roomId string, message []byte) {
	roomsMutex.Lock()
	defer roomsMutex.Unlock()

	for _, conn := range rooms[roomId] {
		conn.WriteMessage(websocket.TextMessage, message)
	}
}
