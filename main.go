package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
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

	r := gin.Default()

	r.POST("/generate", func(c *gin.Context) {
		var requestBody struct {
			Message string `json:"message"`
		}

		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON body"})
			return
		}

		prompt := requestBody.Message
		answer, err := llms.GenerateFromSinglePrompt(ctx, llm, prompt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate response"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"answer": answer})
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
