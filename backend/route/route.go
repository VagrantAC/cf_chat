package route

import (
	"log"
	// "net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)


func NewRoute(engine *gin.Engine) {
  WSRoute(engine)
}

func WSRoute(engine *gin.Engine) {
  group := engine.Group("")
  {
    group.GET("/socket", RunWebsocket)
  }
}

func RunWebsocket(c *gin.Context) {
  
  log.Println("RunWebsocket")
  var upGrader = websocket.Upgrader {
  }
  ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
  if err != nil {
    log.Println("[Websocket] Upgrade error:", err.Error())
    return
  }

  _, message, err := ws.ReadMessage()
  if err != nil {
    log.Println(err.Error())
  }
  ws.WriteMessage(websocket.TextMessage, message)
}
