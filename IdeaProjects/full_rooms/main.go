package main

import (
	"full_call/controllers"
	"full_call/middlewares"
	"full_call/models"
	"github.com/gin-gonic/gin"
)

func main() {

	models.ConnectDataBase()

	r := gin.Default()

	// 방 생성 및 참가/퇴장 라우터 설정
	protected := r.Group("/api")
	protected.Use(middlewares.JwtAuthMiddleware()) // JWT 인증 미들웨어 사용
	protected.POST("/rooms", controllers.CreateRoom)

	r.Run(":8081") // 서버 실행
}
