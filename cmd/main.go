package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"homework_bot/internal/application/services"
	"homework_bot/internal/bot/telegram"
	repository "homework_bot/internal/infrastructure/repositories"
)

var bot *tgbotapi.BotAPI

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	if err := initConfig(); err != nil {
		logrus.Fatalf("init configs err: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("load .env file err: %s", err.Error())
	}

	//db, err := configs.NewPostgresDB(configs.Config{
	//	Host:     viper.GetString("db.host"),
	//	Port:     viper.GetString("db.port"),
	//	Username: viper.GetString("db.username"),
	//	Password: viper.GetString("db.password"),
	//	DBName:   viper.GetString("db.dbname"),
	//	SSLMode:  viper.GetString("db.sslmode"),
	//})

	// Database connection string
	dsn := "root:123456@(8.219.148.240:6033)/gva"

	// Initialize a mysql database connection
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}

	if err != nil {
		logrus.Fatalf("init db err: %s", err.Error())
	}

	//bot, err = tgbotapi.NewBotAPI(os.Getenv("TG_BOT_API"))
	//bot, err = tgbotapi.NewBotAPI("7916934957:AAEy5cOEhSXdAQk5vQyMTVEs8BMRvonm4Ho")
	bot, err = tgbotapi.NewBotAPI("7668068911:AAFOXuA7KpWOfur0rcoVbZTwGOgsBCjkI3s")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Debug = true

	repos := repository.NewRepository(db)
	service := services.NewService(repos)

	tgBot := telegram.NewBot(bot, service)
	err = tgBot.Start()
	if err != nil {
		logrus.Fatalf("bot.start failed: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
