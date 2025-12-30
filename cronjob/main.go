package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/robfig/cron/v3"

	"telegram_bot/constants"
	"telegram_bot/external_services/jokes"
	"telegram_bot/external_services/telegram_bot"
)

var (
	//cronTime = os.Getenv("CRONTIME") //"0 12 * * 3" // Каждый среду (3) в 12:00
	cronTime = "* * * * *"
)

func main() {
	var bot telegram_bot.TgBot
	botToken := constants.EnvBotToken
	chatID, err := strconv.ParseInt(constants.EnvChatID, 10, 0)
	if err != nil {
		log.Fatalf("Не смогли получить переменную окружения ChatID. Ошибка %v", err)
	}
	bot.GetTgBot(botToken)

	// Создаем планировщик
	c := cron.New()

	// Добавляем задачу (каждую неделю в среду в 12:00)
	_, err = c.AddFunc(cronTime, func() {
		bot.SendMsg(chatID, "Ало!Пора бронить будку!")
		joke := jokes.GetAnekdot()
		bot.SendMsg(chatID, fmt.Sprintf("Пока броните вот вам анекдот: %v", joke))
	})

	if err != nil {
		log.Fatal("Ошибка добавления задачи:", err)
	}

	// Запускаем планировщик
	c.Start()

	// Ждем завершения работы
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	// Останавливаем планировщик перед выходом
	c.Stop()
	log.Println("Бот завершает работу")
}
