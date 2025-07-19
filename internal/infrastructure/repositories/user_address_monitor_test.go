package repositories

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"testing"
	"ushield_bot/internal/domain"

	"gorm.io/gorm"
)

func TestUserAddressMonitorRepo_Query(t *testing.T) {
	dsn := "root:1234567890@(156.251.17.226:6033)/gva"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}
	userRepo := NewUserAddressMonitorRepo(db)

	list, errsg := userRepo.Query(context.Background(), 123)
	if errsg != nil {
		panic(errsg)
		return
	}

	for _, v := range list {
		fmt.Printf("%+v\n", v)
	}

}
func TestUserAddressMonitorRepo_Remove(t *testing.T) {
	dsn := "root:1234567890@(156.251.17.226:6033)/gva"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}
	userRepo := NewUserAddressMonitorRepo(db)

	errsg := userRepo.Remove(context.Background(), 123, "111111111111111111111111111111111")
	if errsg != nil {
		panic(errsg)
		return
	}

}
func TestUserAddressMonitorRepo_Create(t *testing.T) {
	dsn := "root:1234567890@(156.251.17.226:6033)/gva"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}
	userRepo := NewUserAddressMonitorRepo(db)
	var record domain.UserAddressMonitor

	record.ChatID = 123
	record.Address = "456789"
	record.Status = 1
	record.Network = "tron"
	errsg := userRepo.Create(context.Background(), &record)
	if errsg != nil {
		panic(errsg)
		return
	}

}
