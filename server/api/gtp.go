package api

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"server/log"
)

type GTPResponse struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type GTPRequest struct {
	ChatId           string       `json:"chatId"`
	Model            string       `json:"model"`
	Stream           bool         `json:"stream"`
	Messages         []ApiMessage `json:"messages"`
	Temperature      float32      `json:"temperature"`
	PresencePenalty  uint8        `json:"presence_penalty"`
	FrequencyPenalty uint8        `json:"frequency_penalty"`
}

func NewGTPRequest() GTPRequest {
	return GTPRequest{
		Stream:           true,
		Model:            "gpt-3.5-turbo",
		Temperature:      0.5,
		PresencePenalty:  0,
		FrequencyPenalty: 0,
	}
}

func (msg GTPRequest) SendMsg(ch map[string]chan []byte) {
	url := "https://chat2.fastgpt.me/api/openai/v1/chat/completions"
	//
	b, err := json.Marshal(msg)
	if err != nil {
		fmt.Println(err)
	}
	//
	r, err := http.Post(url, "application/json", bytes.NewReader(b))

	if err != nil {
		fmt.Println(err)
	}
	defer r.Body.Close()

	if r.StatusCode == 200 {
		reader := bufio.NewReader(r.Body)
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
				ch[msg.ChatId] <- line
			}
		}
	} else {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			log.Log.Errorln(err)
		}
		log.Log.Errorln(string(b))
	}

}
