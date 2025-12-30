package telegram_bot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type TgBot struct {
	bot *tgbotapi.BotAPI
	err error
}

// GetTgBot - создаем апи-клиента телеграм бота (для 1 канала)
func (t *TgBot) GetTgBot(botToken string) {
	if t.bot == nil {
		// Создаем экземпляр бота
		t.bot, t.err = tgbotapi.NewBotAPI(botToken)
		if t.err != nil {
			log.Panic("Ошибка подключения к боту:", t.err)
		}
		log.Println("Бот успешно запущен: @" + t.bot.Self.UserName)
	}
}

func (t *TgBot) SendMsg(chatID int64, message string) {
	msg := tgbotapi.NewMessage(chatID, message)
	if _, err := t.bot.Send(msg); err != nil {
		log.Println("Ошибка отправки сообщения:", err)
	} else {
		log.Println("Сообщение отправлено!")
	}
}
