package util

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
	"github.com/yoyo1025/persona-api/model" // 新しいパッケージをインポート
)

func GenerateAIResponse(persona model.Persona, client *openai.Client) (string, error) {
	prompt := fmt.Sprintf(
		"次のペルソナの情報に基づいて、ペルソナが抱えている問題を解決するアドバイスを提供してください:\n"+
			"名前: %s\n性別: %s\n年齢: %d\n職業: %s\n問題: %s\n行動: %s\n",
		persona.Name, persona.Sex, persona.Age, persona.Profession, persona.Problems, persona.Behavior,
	)

	resp, err := client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model: openai.GPT4,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
	})

	if err != nil {
		return "", fmt.Errorf("OpenAI APIエラー: %v", err)
	}

	return resp.Choices[0].Message.Content, nil
}
