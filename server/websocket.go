package server

import (
	"github.com/gorilla/websocket"
	"github.com/maxcnunes/monitor-api/monitor"
	"log"
	"net/http"
)

// Websocket ...
type Websocket struct {
	data *monitor.DataMonitor
}

type wsConnection struct {
	conn *websocket.Conn
}

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true }, // accepts any origin
	}
)

// SendText ...
func (ws wsConnection) SendText(msg string) {
	log.Printf("Sends ws message [%s]", msg)
	ws.conn.WriteMessage(websocket.TextMessage, []byte(msg))
}

// SendJSON ...
func (ws wsConnection) SendJSON(json interface{}) {
	log.Printf("Sends ws object [%s]", json)
	ws.conn.WriteJSON(json)
}

// Start ...
func (ws Websocket) Start(data *monitor.DataMonitor) func(http.ResponseWriter, *http.Request) {
	log.Print("Starting websocket server")
	ws.data = data

	return func(rw http.ResponseWriter, req *http.Request) {

		conn, err := upgrader.Upgrade(rw, req, nil)
		if err != nil {
			log.Println(err)
			return
		}
		wsConn := &wsConnection{conn: conn}

		log.Printf("WS Connnection entering into the data events loop %s", conn.LocalAddr().String())
		for {
			select {
			case event := <-ws.data.Events:
				switch event.Event {
				case monitor.Added:
					// wsConn.SendText("WS: Added new target")
					wsConn.SendJSON(event)
				case monitor.Removed:
					wsConn.SendJSON(event)
					// wsConn.SendText("WS: Removed old target")
				}
			}
		}
	}
}
