package controllers

import (
	"log"

	"github.com/fi9ish/binanceInfoBot/pkg/binanceRequests"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleCommand(bot *tgbotapi.BotAPI, update tgbotapi.Update, command string, args string) {
	response := switcher(command, args)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
	_, err := bot.Send(msg)
	if err != nil {
		log.Fatal(err)
	}
}

func HandleMessage(bot *tgbotapi.BotAPI, update tgbotapi.Update, message string) {
	response := message
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
	_, err := bot.Send(msg)
	if err != nil {
		log.Fatal(err)
	}
}

func switcher(command string, args string) string {
	switch command {
	case "GetCoins":
		return binanceRequests.GetCoins(args)
	case "GetTickerPrices":
		return binanceRequests.GetTickerPrices(args)
	case "getTickerPriceBySymbol":
		return binanceRequests.GetTickerPriceBySymbol(args)
	default:
		return "Please choose the provided command"
	}
}
