package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/DmKorshenkov/helper/bot/in"
	"github.com/DmKorshenkov/helper/bot/t"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(handler),
	}
	b, err := bot.New(t.Token(), opts...)
	if err != nil {
		// panics for the sake of simplicity.
		// you should handle this error properly in your code.
		print("!")
		panic(err)
	}
	print("now b.Start!!!!\n")
	os.Chdir("DataBase")
	b.Start(ctx)

}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message == nil {
		return
	}
	in.In(update.Message.Text)
}
