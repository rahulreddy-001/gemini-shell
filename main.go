package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"rahul.dev/go/gemini-shell/model"
	"rahul.dev/go/gemini-shell/ui"
)

func handleIntrrupt(ai *model.AI) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	for {
		if os.Interrupt == <-signalChan {
			ai.Close()
			fmt.Println("Exited successfully!")
			os.Exit(0)
		}
	}
}

func main() {
	ctx := context.Background()
	ui.PrintLogo()
	ai := model.NewAI()
	ai.CreateModel(&ctx)
	ai.InitChat()

	go handleIntrrupt(ai)
	for {
		ai.Chat(&ctx)
	}
}
