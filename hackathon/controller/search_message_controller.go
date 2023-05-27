package controller

import (
	"encoding/json"
	"hackathon/model"
	"hackathon/usecase"
	"log"
	"net/http"
)

type SearchMessageController struct {
	SearchMessageUseCase *usecase.SearchMessageUseCase
}

func (c *SearchMessageController) Handle(w http.ResponseWriter, r *http.Request) {
	content := r.URL.Query().Get("content")
	if content == "" {
		log.Println("fail: name is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	messages, err := c.SearchMessageUseCase.Handle(content)
	if err != nil {
		log.Printf("fail: SearchMessageUseCase.Handle, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	messagesRes := make([]model.MessageResForHTTPGet, len(messages))
	for i, u := range messages {
		messagesRes[i] = model.MessageResForHTTPGet{Id: u.Id, Content: u.Content, UserId: u.UserId, ChannelId: u.ChannelId}
	}

	bytes, err := json.Marshal(messagesRes)
	if err != nil {
		log.Printf("fail: json.Marshal, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}
