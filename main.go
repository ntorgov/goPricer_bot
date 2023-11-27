package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	token := os.Getenv("TELEGRAM_BOT_TOKEN")

	if token == "" {
		log.Panic("Переменная окружения TOKEN не установлена")
		return
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

			if update.Message.IsCommand() {
				switch update.Message.Command() {
				case "help":
					msg.Text = "Формат ввода *Вес на ценнике*, через пробел, *Цена за этот вес*"
				case "status":
					msg.Text = "I'm ok."
				default:
					msg.Text = "I don't know that command"
				}

				bot.Send(msg)

				continue
			}

			if !update.Message.IsCommand() {
				command := strings.Split(update.Message.Text, " ")

				if len(command) < 2 {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ошибка конвертирования")
					msg.ReplyToMessageID = update.Message.MessageID
					bot.Send(msg)
					continue
				}

				priceVolume, err := strconv.ParseFloat(command[0], 64)
				if err != nil {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Не понял параметр #1...")
					msg.ReplyToMessageID = update.Message.MessageID
					bot.Send(msg)
					continue
				}

				priceValue, err := strconv.ParseFloat(command[1], 64)
				if err != nil {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Не понял параметр #2...")
					msg.ReplyToMessageID = update.Message.MessageID
					bot.Send(msg)
					continue
				}

				// region Price calculation

				normalPrice := priceValue * 1000 / priceVolume

				log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

				msg = tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Цена за 1000 единиц: %.2f", normalPrice))

				// endregion
			}

			bot.Send(msg)

		}
	}
}
