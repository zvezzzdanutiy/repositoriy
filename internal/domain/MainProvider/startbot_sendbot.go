package MainProvider

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
)

var userID int

const token = "7122476551:AAGRldhloWEs-_jWsEkOTMZEsXhGE0dbXWQ"

func SendAnecdoteToTelegram(anecdote string) error {
	// Формируем URL для отправки сообщения
	endpoint := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token)
	params := url.Values{}
	userIDStr := strconv.Itoa(userID)
	params.Set("chat_id", userIDStr)
	params.Set("text", anecdote)
	_, err := http.PostForm(endpoint, params)
	if err != nil {
		return err
	}

	return nil
}
func StartTelegramBot() {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, _ := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}
		if update.Message.Text == "/start" {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Добро пожаловать! Чтобы получить анекдоты, введите /generate")
			userID = update.Message.From.ID
			bot.Send(msg)
		} else if update.Message.Text == "/generate" {
			go func(update tgbotapi.Update) {
				joke, err := GetJoke()
				fmt.Println("здарова", joke)
				if err != nil {
					log.Println(err)
					return
				}

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, joke)
				bot.Send(msg)
			}(update)
		}
	}
}
func GetJoke() (string, error) {
	err := godotenv.Load()
	if err != nil {
		panic("Can't read env")
	}
	an, err := http.Get(os.Getenv("ANEKDOT_URL"))
	if err != nil {
		fmt.Println(err)
	}
	body, err := io.ReadAll(an.Body)
	if err != nil {
		fmt.Println(err)
	}

	return string(body), err
}
