package util

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
	"github.com/yoyo1025/persona-api/model" // 新しいパッケージをインポート
)

func CreatePersonaFirstComment(persona model.Persona, client *openai.Client) (string, error) {
	prompt := fmt.Sprintf(
		"次の情報は架空の困りごとをペルソナです。あなたはそのペルソナのふりをしてください。ペルソナの設定はその都度付け加えても構いません。次の情報から軽く自己紹介をしてください。改行は要りません。:\n"+
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

func CreateComment(comments []model.Comment, client *openai.Client) (string, error) {
	// 会話履歴を展開
	var conversationHistory string
	for _, comment := range comments {
		conversationHistory += fmt.Sprintf("%s\n", comment.Comment)
	}

	// プロンプトを作成
	prompt := fmt.Sprintf(
		"今から渡す文章のペルソナになり切ったつもりで、次の会話履歴の流れに沿うように簡潔に返答してください。改行は要りません。\n会話履歴:\n%s",
		conversationHistory,
	)

	// OpenAIのChatCompletion APIを呼び出して応答を生成
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
