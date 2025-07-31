package repositories

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"testing"
	"ushield_bot/internal/request"

	"gorm.io/gorm"
)

func TestUserAddressDetectionRepo_GetUserAddressDetectionInfoList(t *testing.T) {
	dsn := "root:1234567890@(156.251.17.226:6033)/ushield"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}
	userRepo := NewUserAddressDetectionRepository(db)

	var info request.UserAddressDetectionSearch
	info.PageInfo.Page = 1
	info.PageInfo.PageSize = 1
	list, total, _ := userRepo.GetUserAddressDetectionInfoList(context.Background(), info, 6620733754)

	fmt.Printf("total:%d\n", total)
	for i := range list {
		fmt.Printf("%+v\n", list[i])
	}

}
