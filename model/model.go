package model

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"

	"rahul.dev/go/gemini-shell/format"
	"rahul.dev/go/gemini-shell/ui"
)

type AI struct {
	client *genai.Client
	model  *genai.GenerativeModel
	cs     *genai.ChatSession
	fs     *format.FormatSelector
}

func NewAI() *AI {
	return &AI{}
}

func (ai *AI) CreateModel(ctx *context.Context) {
	client, err := genai.NewClient(*ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		fmt.Printf("failed to create GenerativeModel client: %v", err)
	}

	ai.client = client
	ai.model = client.GenerativeModel("gemini-pro")
}

func (ai *AI) InitChat() {
	ai.cs = ai.model.StartChat()

	ai.cs.History = []*genai.Content{}
}

func (ai *AI) Chat(ctx *context.Context) {
	prompt := ui.ReadPrompt()
	resp := ai.cs.SendMessageStream(*ctx, genai.Text(prompt))
	ai.printResponse(prompt, resp)
}

func (ai *AI) printResponse(prompt string, iter *genai.GenerateContentResponseIterator) {
	var response string = ""
	ai.fs = format.NewFormatSelector()

	for {
		resp, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		for err, candidate := range resp.Candidates {
			if err != 0 {
				ui.LazyPrint(true, "Error generating candidates", ai.fs)
				return
			}
			for err, part := range candidate.Content.Parts {
				if err != 0 {
					ui.LazyPrint(true, "Error  generating parts", ai.fs)
					return
				}
				partResponse := fmt.Sprintf("%s", part)
				response += partResponse
				ui.LazyPrint(false, partResponse, ai.fs)
			}
		}
	}
	ui.LazyPrint(true, "\n\n", ai.fs)

	ai.saveMessage(prompt, "user")
	ai.saveMessage(response, "model")
}

func (ai *AI) saveMessage(msg string, role string) {
	ai.cs.History = append(ai.cs.History, &genai.Content{
		Parts: []genai.Part{
			genai.Text(msg),
		},
		Role: role,
	})
}

func (ai *AI) Close() {
	ai.client.Close()
	ai.fs.Close()
}
