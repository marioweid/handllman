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
	if err := godotenv.Load("../.env"); err != nil {
		log.Println("Warning: .env file not found, proceeding without it")
	}

	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		log.Fatal("GEMINI_API_KEY is not set in the environment")
	}

	ctx := context.Background()
	llm, err := googleai.New(ctx, googleai.WithAPIKey(apiKey), googleai.WithDefaultModel("gemini-2.0-flash"))
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
	})

	r.POST("/generate", func(c *gin.Context) {
		var requestBody struct {
			Message string `json:"message"`
		}

		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON body"})
			return
		}

		prompt := requestBody.Message
		answer, err := llms.GenerateFromSinglePrompt(c.Request.Context(), llm, prompt)
		if err != nil {
			log.Printf("Error generating response: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate response"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"answer": answer})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default to 8080 if not set
	}
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
