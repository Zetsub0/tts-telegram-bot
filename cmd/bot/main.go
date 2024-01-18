package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/go-yaml/yaml"
	"github.com/zetsub0/tts-telegram-bot/internal/telegram"
	"log"
	"os"
)

type Conf struct {
	Token string `yaml:"token"`
}

var conf Conf

func init() {
	yamlFile, err := os.ReadFile("main.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		panic(err)
	}

}

func main() {
	bot, err := tgbotapi.NewBotAPI(conf.Token)
	if err != nil {
		log.Panic("the api key is not valid")
	}

	//bot.Debug = true
	telegramBot := telegram.NewBot(bot)
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}
}
