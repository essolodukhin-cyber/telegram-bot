package main

import (
	"log"
	"strconv"

	"telegram_bot/constants"
	"telegram_bot/external_services/telegram_bot"
)

func main() {
	var bot telegram_bot.TgBot
	botToken := constants.EnvBotToken
	chatID, err := strconv.ParseInt(constants.EnvChatID, 10, 0)
	if err != nil {
		log.Fatalf("Не смогли получить переменную окружения ChatID. Ошибка %v", err)
	}

	bot.GetTgBot(botToken)
	bot.SendMsg(chatID, "Привет очкошники из папки которую он не должен запускать блеять!")
}
