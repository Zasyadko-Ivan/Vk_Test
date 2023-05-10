package main

import (
	"bufio"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
)

var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("1"),
		tgbotapi.NewKeyboardButton("2"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("4"),
		tgbotapi.NewKeyboardButton("5"),
	),
)

var numericKeyboard2 = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("fmt.Scan(a)"),
		tgbotapi.NewKeyboardButton("a = input()"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("fmt.Scan(&a)"),
		tgbotapi.NewKeyboardButton("fmt.Println(a)"),
	),
)
var numericKeyboard3 = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("10"),
		tgbotapi.NewKeyboardButton("9"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("80"),
		tgbotapi.NewKeyboardButton("8"),
	),
)

var numericKeyboard4 = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Солнце"),
		tgbotapi.NewKeyboardButton("Полярная звезда"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Сириус"),
		tgbotapi.NewKeyboardButton("Альфа Центавра"),
	),
)

func main() {
	bot, err := tgbotapi.NewBotAPI(mustToken())
	if err != nil {
		log.Panicf("Token not found")
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore non-Message updates
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

		switch update.Message.Text {
		case "/start":
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Привет👋\nЭтот бот создан в учебных целях, в нем я создал мини-игру. Чтобы начать игру, напиши слово 'game'")
		case "game", "Game", "GAME":
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Выбери правльный ответ 2+2=")
			msg.ReplyMarkup = numericKeyboard
		case "4":
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Правильный ответ!\n\nСледующий вопрос: Как присвоить значение переменной 'a' в golang?")
			msg.ReplyMarkup = numericKeyboard2
		case "fmt.Scan(&a)":
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Правильный ответ!\n\nСледующий вопрос: Сколько ног у паука?")
			msg.ReplyMarkup = numericKeyboard3
		case "8":
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Правильный ответ!\n\nИ последний вопрос: Какая зведа ближе всего расположена к земле?")
			msg.ReplyMarkup = numericKeyboard4
		case "Солнце":
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Правильный ответ!\n\nCпасибо за уделенное время☺️")
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		case "1", "2", "5", "fmt.Scan(a)", "a = input()", "fmt.Println(a)", "10", "9", "80", "Полярная звезда", "Сириус", "Альфа Центавра":
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Это не правильный ответ, попробуй еще \U0001FAE3")
		default:
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Для того, чтобы начать игру введи слово 'game'")
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}

func mustToken() string {
	file, err := os.Open("tokenTelegram")
	// Необходим файл tokenTelegram с токином от телеграмм бота
	if err != nil {
		log.Panicf("File tokenTelegram not found")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	return scanner.Text()
}
