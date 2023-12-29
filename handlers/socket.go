package handler

import (
	"encoding/json"
	"fmt"

	model "github.com/Faizal-Asep/test-tech/models"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var Upgrader = websocket.Upgrader{}
var WsConn []*websocket.Conn

func init() {
	model.Notif = make(chan model.Notification)
	go func() {
		for {
			data := <-model.Notif
			b, _ := json.Marshal(data)
			for idx, con := range WsConn {
				if err := con.WriteMessage(websocket.TextMessage, b); err != nil {
					con.Close()
					WsConn = remove(WsConn, idx)
				}
			}
		}
	}()
}
func remove(slice []*websocket.Conn, s int) []*websocket.Conn {
	return append(slice[:s], slice[s+1:]...)
}

func Socket(c *gin.Context) {

	conn, err := Upgrader.Upgrade(c.Writer, c.Request, nil)
	WsConn = append(WsConn, conn)
	if err != nil {
		fmt.Println(err)
		return
	}
}
