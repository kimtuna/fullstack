package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"log"
)

var DB *gorm.DB

func ConnectDataBase() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// MySQL 관련 환경 변수 설정
	Dbdriver := "mysql"     // 드라이버를 mysql로 변경
	DbHost := "localhost"   // 호스트 이름 (도커와 같은 환경에서는 서비스 이름을 사용)
	DbUser := "tuna"        // 사용자 이름
	DbPassword := "tuna123" // 비밀번호
	DbName := "fulldb"      // 데이터베이스 이름
	DbPort := "3306"        // 포트 번호

	// DB URL 생성
	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)

	DB, err = gorm.Open(Dbdriver, DBURL)

	if err != nil {
		fmt.Println("Cannot connect to database ", Dbdriver)
		log.Fatal("connection error:", err)
	} else {
		fmt.Println("We are connected to the database ", Dbdriver)
	}

	DB.AutoMigrate(&Room{}) // User 모델에 맞게 테이블을 자동으로 마이그레이션합니다.
}
