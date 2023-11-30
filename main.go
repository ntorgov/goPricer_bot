package main

import (
	"Pricer/internal"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
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
				case "price":
					msg.Text = "Формат ввода *Вес на ценнике*, через пробел, *Цена за этот вес*"
				default:
					msg.Text = "I don't know that command"
				}

				bot.Send(msg)

				continue
			}

			if !update.Message.IsCommand() {
				priceVolume, priceValue, err := internal.CommandParser(update.Message.Text) // strings.Split(update.Message.Text, " ")

				if err != nil {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Не понял...")
					msg.ReplyToMessageID = update.Message.MessageID
					bot.Send(msg)
					continue
				}

				// region Price calculation

				normalPrice := internal.CalculateCleanPrice(priceVolume, priceValue)

				log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

				msg = tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Цена за 1000 единиц: %.2f", normalPrice))

				// endregion
			}

			bot.Send(msg)

		}
	}
}
