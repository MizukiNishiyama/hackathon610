package controller

import (
	"encoding/json"
	"hackathon/model"
	"hackathon/usecase"
	"io"
	"log"
	"net/http"
)

type RegisterChannelController struct {
	RegisterChannelUseCase *usecase.RegisterChannelUseCase
}

func (c *RegisterChannelController) Handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("fail: io.ReadAll, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var channelReq model.ChannelReqForHTTPPost
	if err := json.Unmarshal(body, &channelReq); err != nil {
		log.Printf("fail: json.Unmarshal, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	channel, err := c.RegisterChannelUseCase.Handle(channelReq)
	if err != nil {
		log.Printf("fail: RegisterUserUseCase.Handle, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	channelRes := model.ChannelResForHTTPPost{Id: channel.Id}
	bytes, err := json.Marshal(channelRes)
	if err != nil {
		log.Printf("fail: json.Marshal, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	w.Write(bytes)
}
