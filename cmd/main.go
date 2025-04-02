package main

import (
	"bytes"
	"encoding/json"
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
	"net/http"
	"time"
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
	dsn := "root:12345678901234567890@(156.251.17.226:6033)/gva"

	// Initialize a mysql database connection
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}

	if err != nil {
		logrus.Fatalf("init db err: %s", err.Error())
	}

	//bot, err = tgbotapi.NewBotAPI(os.Getenv("TG_BOT_API"))
	const token = "7551982200:AAHdSLHtqDj25ugn3uD1hth9i2iRU8OYWnU"
	bot, err = tgbotapi.NewBotAPI(token)
	//bot, err = tgbotapi.NewBotAPI("7668068911:AAFOXuA7KpWOfur0rcoVbZTwGOgsBCjkI3s")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Debug = true

	repos := repository.NewRepository(db)
	service := services.NewService(repos)

	tgBot := telegram.NewBot(bot, service)
	go asyncNotify(tgBot, token)

	err = tgBot.Start()

	if err != nil {
		logrus.Fatalf("bot.start failed: %s", err.Error())
	}
}

func asyncNotify(tgBot *telegram.Bot, token string) {
	for {
		fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>Hello, World<<<<<<<<<<<<<<<<<<<<<<<<<<<<")
		addresses, _ := tgBot.GetServices().IUserService.NotifyTronAddress()

		for _, address := range addresses {

			notify(address.Associates, token, address.TronAddress)
			tgBot.GetServices().IUserService.DisableTronAddress(address.TronAddress)

		}

		eth_addresses, _ := tgBot.GetServices().IUserService.NotifyEthereumAddress()

		for _, address := range eth_addresses {
			notify(address.Associates, token, address.EthAddress)
			tgBot.GetServices().IUserService.DisableTronAddress(address.EthAddress)

		}

		time.Sleep(60 * time.Second) // 等待 30秒
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func notify(_chatID string, _token string, _address string) {
	message := map[string]string{
		"chat_id": _chatID, // 或直接用 chat_id 如 "123456789"=
		"text":    "‼️‼️" + _address + "请注意地址即将被拉入黑名单" + "️‼️‼️️",
	}
	// 转换为 JSON
	jsonData, err := json.Marshal(message)
	if err != nil {
		fmt.Println("JSON 编码失败:", err)
		return
	}

	// 发送 POST 请求到 Telegram Bot API
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", _token)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("发送消息失败:", err)
		return
	}
	defer resp.Body.Close()

	// 打印响应结果
	fmt.Println("消息发送状态:", resp.Status)
}
