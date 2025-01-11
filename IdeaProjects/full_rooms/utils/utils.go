package utils

import (
	"fmt"
)

func GenerateRoomURL(roomID string, userID uint) string {
	baseURL := "http://yourdomain.com/rooms" // 환경변수처리
	return fmt.Sprintf("%s/%s?user_id=%d", baseURL, roomID, userID)
}
