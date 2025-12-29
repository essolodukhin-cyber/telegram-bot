package main

import (
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/robfig/cron/v3"
)

func main() {
	// Токен бота
	botToken := os.Getenv("BOTTOKEN")
	//botToken := "8437986859:AAHRZcu3gc3kn5cHAoV1Xn2WtHrylxIT9p4" // Например: "8206421086:AAHysZGoynV-vDc6hepIfV1zLp6UvNh9t3I"
	chatID, err := strconv.ParseInt(os.Getenv("CHATID"), 10, 0)
	if err != nil {
		log.Fatal(err)
	}
	//chatID := int64(-1002880878621) // ID чата (число, например, -925488365)

	// Создаем экземпляр бота
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic("Ошибка подключения к боту:", err)
	}
	log.Println("Бот успешно запущен: @" + bot.Self.UserName)

	msg := tgbotapi.NewMessage(chatID, "Привет очкошники!")
	if _, err := bot.Send(msg); err != nil {
		log.Println("Ошибка отправки сообщения:", err)
	} else {
		log.Println("Сообщение отправлено!")
	}

	// Создаем планировщик
	c := cron.New()

	// Добавляем задачу (каждую неделю в среду в 12:00)
	_, err = c.AddFunc("0 12 * * 3", func() { // Каждый среду (3) в 12:00
		msg := tgbotapi.NewMessage(chatID, "Ало!Пора бронить будку!")
		if _, err := bot.Send(msg); err != nil {
			log.Println("Ошибка отправки сообщения:", err)
		} else {
			log.Println("Сообщение отправлено!")
		}
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
