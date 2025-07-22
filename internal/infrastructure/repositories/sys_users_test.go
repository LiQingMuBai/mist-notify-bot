package repositories

import (
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestSysUsersRepository_Find(t *testing.T) {
	dsn := "root:1234567890@(156.251.17.226:6033)/ushield"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}
	userRepo := NewSysUsersRepository(db)

	address, depositAddress, errsg := userRepo.Find(context.Background(), "admin")
	if errsg != nil {
		panic(errsg)
		return
	}

	fmt.Printf("address:%v\n", address)
	fmt.Printf("depositAddress:%v\n", depositAddress)

}
