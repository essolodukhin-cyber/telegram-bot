package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/robfig/cron/v3"
)

var (
	envBotToken = os.Getenv("BOTTOKEN")
	envChatID   = os.Getenv("CHATID")
	cronTime    = os.Getenv("CRONTIME") //"0 12 * * 3" // Каждый среду (3) в 12:00
)

func main() {
	// Токен бота
	botToken := envBotToken
	//botToken := "8437986859:AAHRZcu3gc3kn5cHAoV1Xn2WtHrylxIT9p4"
	chatID, err := strconv.ParseInt(envChatID, 10, 0)
	if err != nil {
		log.Fatalf("Не смогли получить переменную окружения ChatID. Ошибка %v", err)
	}
	//chatID := int64(-1002880878621)

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
	_, err = c.AddFunc(cronTime, func() {
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

	resp, err := http.Get("http://rzhunemogu.ru/RandJSON.aspx?CType=11")
	if err != nil {
		log.Fatalf("Не смогли получить шутку. Ошибка %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении:", err)
		return
	}

	// Простая обработка ответа без использования json.Unmarshal
	content := string(body)
	content = strings.TrimPrefix(content, `{"content":"`)
	content = strings.TrimSuffix(content, `"}`)
	content = strings.ReplaceAll(content, `\r\n`, "\n")

	fmt.Println(content)

	anekdot := tgbotapi.NewMessage(chatID, fmt.Sprintf("Пока броните вот вам анекдот: %v", content))
	if _, err = bot.Send(anekdot); err != nil {
		log.Println("Ошибка отправки анекдота:", err)
	} else {
		log.Println("Анекдот отправлен!")
	}
}
