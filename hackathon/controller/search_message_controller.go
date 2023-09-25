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
	channelid := r.URL.Query().Get("channelid")
	if channelid == "" {
		log.Println("fail: channelid is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	messages, err := c.SearchMessageUseCase.Handle(channelid)
	if err != nil {
		log.Printf("fail: SearchMessageUseCase.Handle, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	messagesRes := make([]model.MessageResForHTTPGet, len(messages))
	for i, u := range messages {
		messagesRes[i] = model.MessageResForHTTPGet{Id: u.Id, Content: u.Content, UserId: u.UserId, ChannelId: u.ChannelId, Time: u.Time}
	}

	bytes, err := json.Marshal(messagesRes)
	if err != nil {
		log.Printf("fail: json.Marshal, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Write(bytes)
}
