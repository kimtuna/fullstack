package middlewares

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 요청 헤더에서 토큰 추출
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token not provided"})
			c.Abort()
			return
		}

		// Bearer 접두사 제거
		if len(token) > 7 && token[:7] == "Bearer " {
			token = token[7:] // "Bearer " 부분 제거
		}

		// 다른 서비스에 인증 요청
		authURL := "http://localhost:8080/api/admin/auth"
		log.Printf("call: jwtmiddleware token test: %s", token)

		// POST 요청 생성 (토큰을 본문이 아닌 헤더에 전달)
		req, err := http.NewRequest("POST", authURL, nil) // 본문 없이 요청 생성
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
			c.Abort()
			return
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+token) // 헤더에 토큰 추가

		// 요청 실행
		resp, err := http.DefaultClient.Do(req)
		if err != nil || resp.StatusCode != http.StatusOK {
			c.JSON(http.StatusUnauthorized, gin.H{"call": "Unauthorized"})
			c.Abort()
			return
		}
		defer resp.Body.Close()

		// 토큰이 유효함을 의미하는 응답을 받았을 때
		c.Next()
	}
}
