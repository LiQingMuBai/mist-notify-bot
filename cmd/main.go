package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"

	"github.com/joho/godotenv"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"time"
	"ushield_bot/internal/application/services"
	"ushield_bot/internal/bot/telegram"
	repository "ushield_bot/internal/infrastructure/repositories"
)

//机器人完成
//同步任务监听地址充值
//调用trxfee平台发能量
//地址监控

var bot *tgbotapi.BotAPI

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	if err := initConfig(); err != nil {
		logrus.Fatalf("init configs err: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("load .env file err: %s", err.Error())
	}

	// Database connection string
	host := viper.GetString("db.host")
	port := viper.GetString("db.port")
	username := viper.GetString("db.username")
	password := viper.GetString("db.password")
	dbname := viper.GetString("db.dbname")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}

	if err != nil {
		logrus.Fatalf("init db err: %s", err.Error())
	}

	token := os.Getenv("TG_BOT_API")
	_cookie := os.Getenv("COOKIE")
	bot, err = tgbotapi.NewBotAPI(token)

	//bot, err = tgbotapi.NewBotAPI(token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Debug = true

	repos := repository.NewRepository(db)
	service := services.NewService(repos)

	agent := viper.GetString("agent")

	log.Println("agent:", agent)
	tgBot := telegram.NewBot(bot, service, _cookie, agent, db)
	//go asyncNotify(tgBot, token)

	err = tgBot.Start()

	if err != nil {
		logrus.Fatalf("bot.start failed: %s", err.Error())
	}
}

func asyncNotify(tgBot *telegram.Bot, token string) {
	for {

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
