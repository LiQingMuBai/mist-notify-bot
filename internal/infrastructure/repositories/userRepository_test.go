package repositories

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	//_ "github.com/lib/pq"
	"log"
	"testing"
	"ushield_bot/internal/domain"
)

func TestUserRepository_UpdateAddress(t *testing.T) {
	dsn := "root:12345678901234567890@(156.251.17.226:6033)/gva"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}
	//userRepo := NewUserRepository(db)

	userRepo := NewUserRepository(db)
	user, err := userRepo.GetByUserID("7347235462")
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("%#v\n", user)

	log.Println("=============================================================")
	updateUser := domain.User{
		Id:       3,
		Username: "avachow101",
		Key:      "6cec7800bca14e2f28d44e731a437e991399e1410973b02b74eb8217b04a1f96",
		Address:  "TJLmSN4sbsAg4bKxqSv9SZL1RqnWFRfrRm",
	}
	err = userRepo.UpdateAddress(updateUser)
	if err != nil {
		log.Println(err)
	}
}
func TestUserRepository_GetByUsername(t *testing.T) {
	dsn := "root:12345678901234567890@(156.251.17.226:6033)/gva"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}
	//userRepo := NewUserRepository(db)

	userRepo := NewUserRepository(db)

	err = userRepo.Create(domain.User{
		Username: "masion",
		UserID:   "11223",
	})

	if err != nil {
		log.Println(err)
	}
}
func TestUserRepository_GetByUserID(t *testing.T) {
	dsn := "root:12345678901234567890@(156.251.17.226:6033)/gva"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}
	//userRepo := NewUserRepository(db)

	userRepo := NewUserRepository(db)

	userRepo.FetchNewestAddress()
}
