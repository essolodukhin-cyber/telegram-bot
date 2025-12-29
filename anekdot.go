package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

func GetAnekdot() string {
	resp, err := http.Get("http://rzhunemogu.ru/RandJSON.aspx?CType=11")
	if err != nil {
		log.Fatalf("Не смогли получить шутку. Ошибка %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении:", err)
	}

	// Конвертируем из Windows-1251 в UTF-8
	decoder := charmap.Windows1251.NewDecoder()
	utf8Body, _, err := transform.Bytes(decoder, body)
	if err != nil {
		fmt.Printf("Ошибка при конвертации кодировки: %v\n", err)
	}

	content := string(utf8Body)
	content = strings.TrimPrefix(content, `{"content":"`)
	content = strings.TrimSuffix(content, `"}`)
	content = strings.ReplaceAll(content, `\r\n`, "\n") // Убираем экранированные символы

	return content
}
