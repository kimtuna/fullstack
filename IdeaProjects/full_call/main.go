package main

import (
  "github.com/gin-gonic/gin"
  "github.com/gorilla/websocket"
  "log"
)

var upgrader = websocket.Upgrader{}

func main() {
	r := gin.Default()

	r.GET("/ws", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Println("Error during connection upgrade:", err)
			return
		}
		defer conn.Close()

		for {
			messageType, msg, err := conn.ReadMessage()
			if err != nil {
				log.Println("Error reading message:", err)
				break
			}

			// 모든 클라이언트에 메시지 전송
			// 이 부분에서 신호 처리를 구현
			err = conn.WriteMessage(messageType, msg)
			if err != nil {
				log.Println("Error writing message:", err)
				break
			}
		}
	})

	r.Run(":8080") // 서버 실행
}
