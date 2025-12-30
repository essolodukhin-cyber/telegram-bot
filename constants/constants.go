package constants

import (
	"os"
)

// переменные окружения
var (
	EnvBotToken = os.Getenv("BOTTOKEN")
	EnvChatID   = os.Getenv("CHATID")
	CronTime    = os.Getenv("CRONTIME") //"0 12 * * 3" // Каждый среду (3) в 12:00
)
