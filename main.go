package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/googleai"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	ctx := context.Background()
	apiKey := os.Getenv("GEMINI_API_KEY")
	llm, err := googleai.New(ctx, googleai.WithAPIKey(apiKey), googleai.WithDefaultModel("gemini-2.0-flash"))
	if err != nil {
		log.Fatal(err)
	}

	prompt := "What is 2 + 2? Asnwer very short!"
	answer, err := llms.GenerateFromSinglePrompt(ctx, llm, prompt)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(answer)
}
