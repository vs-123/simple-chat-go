package main

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	m := melody.New()
	// port := os.Getenv("PORT")

	r.GET("/", func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, "index.html")
	})

	r.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		if splitMsg := strings.Split(string(msg), " "); splitMsg[0] == "CLIENTCLOSING" {
			m.Broadcast([]byte("<i>" + splitMsg[1] + " <red>left</red> the chat</i>"))
		} else {
			m.Broadcast(msg)
		}
	})

	r.Run(":" + os.Getenv("PORT"))
}
