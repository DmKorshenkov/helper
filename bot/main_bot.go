package main

import (
	"context"
	"fmt"
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
	err = os.Chdir("../")
	if err != nil {
		fmt.Println("!")
	}
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
	helper()
	in.In(update.Message.Text)
}

func helper() {
	f, err := os.Create("test.json")
	if err != nil {
		fmt.Println("s")
	}
	f.Close()
}
