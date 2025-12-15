package config

import (
    "context"
    "log"
    "os"

    "google.golang.org/genai"
)

var (
    GeminiClient *genai.Client
    GeminiConfig *genai.GenerateContentConfig
)

func InitGemini() {
    ctx := context.Background()

    apiKey := os.Getenv("GEMINI_API_KEY")
    if apiKey == "" {
        log.Fatal("GEMINI_API_KEY is empty")
    }

    client, err := genai.NewClient(ctx, &genai.ClientConfig{
        APIKey: apiKey,
    })
    if err != nil {
        log.Fatal(err)
    }

    GeminiClient = client

    // System prompt global
    GeminiConfig = &genai.GenerateContentConfig{
        SystemInstruction: genai.NewContentFromText(
            "Kamu adalah guru coding profesional",
            genai.RoleUser,
        ),
    }
}
