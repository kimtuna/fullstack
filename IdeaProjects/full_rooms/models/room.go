package models

type Room struct {
	ID       string `gorm:"primaryKey"`
	ParentID uint   // 사용자 UID
	Name     string // 방 이름
}
