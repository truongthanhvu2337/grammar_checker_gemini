package pkg

import (
	"context"
	"fmt"
	"log"

	"grammar-checker/pkg/config"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type GeminiClient struct {
	client *genai.Client
	model  *genai.GenerativeModel
}

func LoadAPIKey() string {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("⚠️ Không tìm thấy file .env, sử dụng biến môi trường hệ thống")
	}

	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		fmt.Println("❌ Chưa có GEMINI_API_KEY, vui lòng kiểm tra file .env")
		os.Exit(1)
	}

	return apiKey
}

// NewGeminiClient tạo một client mới
func NewGeminiClient() *GeminiClient {
	ctx := context.Background()
	apiKey := config.LoadAPIKey()

	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("❌ Lỗi tạo Gemini client: %v", err)
	}

	model := client.GenerativeModel("gemini-2.0-flash")
	model.SetTemperature(1)
	model.SetTopK(40)
	model.SetTopP(0.95)
	model.SetMaxOutputTokens(8192)
	model.ResponseMIMEType = "text/plain"

	return &GeminiClient{client: client, model: model}
}

func (g *GeminiClient) CheckGrammar(inputText string) (string, error) {
	ctx := context.Background()
	session := g.model.StartChat()

	resp, err := session.SendMessage(ctx, genai.Text(fmt.Sprintf("Hãy sửa lỗi ngữ pháp của câu sau: %s", inputText)))
	if err != nil {
		return "", fmt.Errorf("❌ Lỗi khi gửi tin nhắn đến Gemini API: %v", err)
	}

	return resp.Candidates[0].Content.Parts[0].String(), nil
}

func (g *GeminiClient) Close() {
	g.client.Close()
}
