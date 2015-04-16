package server

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/maxcnunes/go-uptime-api/monitor/data"
	"github.com/maxcnunes/go-uptime-api/monitor/entities"
)

// Websocket aggregates all the web socket configuration and actions
type Websocket struct {
	data *data.DataMonitor
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
	log.Printf("Sends ws object [%v]", json)
	if err := ws.conn.WriteJSON(json); err != nil {
		log.Printf(" Error sending ws object [%v]: %v", json, err)
	}
}

// Start ...
func (ws Websocket) Start(dm *data.DataMonitor) func(http.ResponseWriter, *http.Request) {
	log.Print("Starting websocket server")
	ws.data = dm

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
				case entities.Added:
					// wsConn.SendText("WS: Added new target")
					wsConn.SendJSON(event)
				case entities.Removed:
					wsConn.SendJSON(event)
					// wsConn.SendText("WS: Removed old target")
				}
			}
		}
	}
}
