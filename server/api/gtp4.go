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

type GTP4Response struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type MSG struct {
	Content     string `json:"content"`
	Id          string `json:"id"`
	Role        string `json:"role"`
	Timestamp   int64  `json:"timestamp"`
	Who         string `json:"who"`
	IsStreaming bool   `json:"isStreaming"`
}

type GTP4Request struct {
	Id         string `json:"id"`
	BotId      string `json:"botId"`
	Session    string `json:"session"`
	ClientId   string `json:"clientId"`
	ContextId  string `json:"contextId"`
	NewMessage string `json:"newMessage"`
	Stream     bool   `json:"stream"`
	Messages   []MSG  `json:"messages"`
}

func NewGTP4Request() GTP4Request {
	return GTP4Request{
		Id:      "default",
		BotId:   "default",
		Session: "N/A",
		Stream:  true,
	}
}

func (msg GTP4Request) SendMsg(ch map[string]chan []byte) {
	url := "https://free-chatgpt.online/wp-json/mwai-ui/v1/chats/submit"
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
				var res GTP4Response
				if err := json.Unmarshal(line[6:], &res); err != nil {
					log.Log.Errorln(err)
				}
				// if res.Type == "live" {

				// 	fmt.Printf("%s", res.Data)
				// }
				// if res.Type == "end" {
				// 	fmt.Print("\n")
				// 	fmt.Println(res)
				// }
				ch[msg.ContextId] <- line

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
