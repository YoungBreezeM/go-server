package api

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"server/config"
	"server/log"

	http "github.com/bogdanfinn/fhttp"
	tls_client "github.com/bogdanfinn/tls-client"
	"github.com/google/uuid"
)

type ApiMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
type APIRequest struct {
	ChatId          string       `json:"chatId"`
	Messages        []ApiMessage `json:"messages"`
	Model           string       `json:"model"`
	ParentMessageID string       `json:"parent_message_id,omitempty"`
	ConversationID  string       `json:"conversation_id,omitempty"`
}

type chatgpt_message struct {
	ID      uuid.UUID       `json:"id"`
	Author  chatgpt_author  `json:"author"`
	Content chatgpt_content `json:"content"`
}

type chatgpt_content struct {
	ContentType string   `json:"content_type"`
	Parts       []string `json:"parts"`
}

type chatgpt_author struct {
	Role string `json:"role"`
}

type ChatGPTRequest struct {
	Action                     string            `json:"action"`
	Messages                   []chatgpt_message `json:"messages"`
	ParentMessageID            string            `json:"parent_message_id,omitempty"`
	ConversationID             string            `json:"conversation_id,omitempty"`
	Model                      string            `json:"model"`
	HistoryAndTrainingDisabled bool              `json:"history_and_training_disabled"`
}

type ChatGPTResponse struct {
	Message        Message     `json:"message"`
	ConversationID string      `json:"conversation_id"`
	Error          interface{} `json:"error"`
}

type Message struct {
	ID         string      `json:"id"`
	Author     Author      `json:"author"`
	CreateTime float64     `json:"create_time"`
	UpdateTime interface{} `json:"update_time"`
	Content    Content     `json:"content"`
	EndTurn    interface{} `json:"end_turn"`
	Weight     float64     `json:"weight"`
	Metadata   Metadata    `json:"metadata"`
	Recipient  string      `json:"recipient"`
}

type Content struct {
	ContentType string   `json:"content_type"`
	Parts       []string `json:"parts"`
}

type Author struct {
	Role     string                 `json:"role"`
	Name     interface{}            `json:"name"`
	Metadata map[string]interface{} `json:"metadata"`
}

type Metadata struct {
	Timestamp     string         `json:"timestamp_"`
	MessageType   string         `json:"message_type"`
	FinishDetails *FinishDetails `json:"finish_details"`
	ModelSlug     string         `json:"model_slug"`
	Recipient     string         `json:"recipient"`
}

type FinishDetails struct {
	Type string `json:"type"`
	Stop string `json:"stop"`
}

func NewChatGPTRequest() ChatGPTRequest {
	return ChatGPTRequest{
		Action:                     "next",
		ParentMessageID:            uuid.NewString(),
		Model:                      "text-davinci-002-render-sha",
		HistoryAndTrainingDisabled: false,
	}
}

func (c *ChatGPTRequest) AddMessage(role string, content string) {
	c.Messages = append(c.Messages, chatgpt_message{
		ID:      uuid.New(),
		Author:  chatgpt_author{Role: role},
		Content: chatgpt_content{ContentType: "text", Parts: []string{content}},
	})
}

var (
	jar     = tls_client.NewCookieJar()
	options = []tls_client.HttpClientOption{
		tls_client.WithTimeoutSeconds(360),
		tls_client.WithClientProfile(tls_client.Firefox_110),
		tls_client.WithNotFollowRedirects(),
		tls_client.WithCookieJar(jar), // create cookieJar instance and pass it as argument
		// Disable SSL verification
		tls_client.WithInsecureSkipVerify(),
	}
	client, _ = tls_client.NewHttpClient(tls_client.NewNoopLogger(), options...)
)

func (msg APIRequest) SendMsg(ch map[string]chan []byte) {
	err := client.SetProxy("socks5://127.0.0.1:8088")
	if err != nil {
		log.Log.Errorln(err)
		return
	}
	apiUrl := "https://chat.openai.com/backend-api/conversation"
	//
	chatReq := NewChatGPTRequest()
	if msg.ConversationID != "" && msg.ParentMessageID != "" {
		chatReq.ConversationID = msg.ConversationID
		chatReq.ParentMessageID = msg.ParentMessageID
	}

	chatReq.AddMessage(msg.Messages[0].Role, msg.Messages[0].Content)
	//
	body_json, err := json.Marshal(chatReq)
	if err != nil {
		log.Log.Errorln(err)
	}

	request, err := http.NewRequest(http.MethodPost, apiUrl, bytes.NewBuffer(body_json))
	if err != nil {
		log.Log.Errorln(err)
	}
	// Clear cookies
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36")
	request.Header.Set("Accept", "*/*")
	request.Header.Set("Authorization", "Bearer "+config.ACCESSTOKEN)
	request.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	//
	response, err := client.Do(request)
	if err != nil {
		log.Log.Errorln(err)
	}
	defer response.Body.Close()
	log.Log.Info(response.Status)
	//
	if response.StatusCode == 200 {
		reader := bufio.NewReader(response.Body)
		for {
			line, err := reader.ReadBytes('\n')
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Log.Errorln("Failed to read response:", err)
				return
			}
			//
			if len(line) > 1 {
				fmt.Println(msg.ChatId)
			}
		}

	} else {
		log.Log.Warning(response.Status)
	}

}
