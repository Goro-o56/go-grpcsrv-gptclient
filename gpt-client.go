package main

import (
	"context"
	"fmt"
	"github.com/PullRequestInc/go-gpt3"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GptClient(prompt []string) string {
	godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	ctx := context.Background()
	client := gpt3.NewClient(apiKey)
	gpt3.WithDefaultEngine(gpt3.TextDavinci003Engine)

	var stop []string = []string{"\n"}
	resp, err := client.Completion(ctx, gpt3.CompletionRequest{
		Prompt:           prompt,
		MaxTokens:        gpt3.IntPtr(100),
		Temperature:      gpt3.Float32Ptr(0),
		TopP:             gpt3.Float32Ptr(0),
		PresencePenalty:  0,
		FrequencyPenalty: 0,
		Stop:             stop,
	})

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(resp.Choices[0].Text)

	return resp.Choices[0].Text
}
