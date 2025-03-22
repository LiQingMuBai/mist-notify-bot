package repositories

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"homework_bot/internal/domain"
	"log"
	"testing"
)

var db *sqlx.DB

//db:
//host: "8.219.148.240"
//port: "5432"
//username: "admin"
//password: "severn_2025"
//dbname: "mydb"
//sslmode: "disable"

//##docker run --name postgresql -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=severn_2025 -e POSTGRES_DB=mydb -p 5432:5432 -d postgres

func TestUserRepository_GetByUsername(t *testing.T) {

	//connect to a PostgreSQL database
	// Replace the connection details (user, dbname, password, host) with your own
	db, err := sqlx.Connect("postgres", "user=admin dbname=mydb sslmode=disable password=severn_2025 host=8.219.148.240")
	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	// Test the connection to the database
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully Connected")
	}

	place := domain.User{}

	rows, _ := db.Queryx("SELECT  id,username,amount,associates, tron_amount,tron_address,eth_address,eth_amount FROM tg_users")
	for rows.Next() {
		err := rows.StructScan(&place)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf(">>>>>>>>>>>>>>>>>>>>>>name:%s \n", place.Username)

	}

	log.Printf("%#v\n", place)

	jason := domain.User{}
	err = db.Get(&jason, "SELECT  id,username,amount,associates, tron_amount,tron_address,eth_address,eth_amount ,create_at,update_at FROM tg_users WHERE username=$1", "avachow101")
	fmt.Printf("%#v\n", jason.Id.String())
	fmt.Printf("%#v\n", jason.CreatedAt.String())

	//
	//db, err := sqlx.Connect("postgres", "user=foo dbname=bar sslmode=disable")
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//query := "SELECT id,associates,amount, tron_amount,tron_address,eth_address,eth_amount FROM tg_users WHERE username = ? "
	//var user domain.User
	//err = db.Get(&user, query, "avachow101")
	//if err != nil {
	//	fmt.Printf("get failed, err:%v\n", err)
	//	return
	//}
	//fmt.Printf(" name:%s \n", user.Username)
}
