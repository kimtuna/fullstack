package controllers

import (
	"bytes"
	"encoding/json"
	"full_call/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
)

// CreateRoom 핸들러
func CreateRoom(c *gin.Context) {
	// 요청 헤더에서 토큰 추출
	token := c.Request.Header.Get("Authorization")

	log.Printf("createroom token test %s", token)
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token not provided"})
		return
	}

	//// Bearer 접두사 제거
	//if len(token) > 7 && token[:7] == "Bearer " {
	//	token = token[7:] // "Bearer " 부분 제거
	//}

	// UID를 가져오기 위해 /api/admin/auth 호출
	authURL := "http://localhost:8080/api/admin/auth"
	reqBody, err := json.Marshal(gin.H{"token": token})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request body"})
		return
	}

	// 요청 생성
	req, err := http.NewRequest("POST", authURL, bytes.NewBuffer(reqBody))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token) // 토큰을 헤더에 추가

	// 요청 실행
	resp, err := http.DefaultClient.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusUnauthorized, gin.H{"call": "Unauthorized"})
		return
	}
	defer resp.Body.Close()

	// UID 추출
	var authResponse struct {
		UID uint `json:"uid"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&authResponse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode UID"})
		return
	}

	uid := authResponse.UID // UID 저장

	// 방 생성 데이터 구조
	var roomData struct {
		RoomName string `json:"roomName"`
	}
	if err := c.ShouldBindJSON(&roomData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// 방 ID 생성 (예: UUID)
	roomID := generateRoomID() // UUID 생성 함수

	// 데이터베이스에 방 정보 저장
	err = saveRoomToDatabase(roomID, uid, roomData.RoomName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create room"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Room created successfully", "roomID": roomID, "uid": uid})
}

// 데이터베이스에 방 정보를 저장하는 함수
func saveRoomToDatabase(roomID string, uid uint, roomName string) error {
	room := models.Room{ID: roomID, ParentID: uid, Name: roomName} // Room 모델 사용
	result := models.DB.Create(&room)                              // db.DB 사용
	return result.Error
}

// 방 ID 생성 함수
func generateRoomID() string {
	return uuid.New().String() // UUID 생성
}
