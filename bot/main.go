package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	os.Chdir("./DataBase")

	opts := []bot.Option{
		bot.WithDefaultHandler(handler),
	}
	os.Setenv("token", "7962376260:AAElZgEDL9BHRBeozjqsRG0fvvGn6ftqERc")

	b, err := bot.New(os.Getenv("token"), opts...)
	if err != nil {
		// panics for the sake of simplicity.
		// you should handle this error properly in your code.
		print("!")
		panic(err)
	}
	print("now b.Start!!!!\n")
	b.Start(ctx)

}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message == nil {
		return
	}

	help(update.Message.Text)
}

func help(str string) {
	f, _ := os.OpenFile("tmp.txt", os.O_RDWR, 0666)

	_, err := f.WriteString(str)
	if err != nil {
		log.Println(err.Error())
	}
	f.Close()
}
