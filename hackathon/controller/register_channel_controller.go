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
	w.Header().Set("Content-Type", "application/json")

	w.Write(bytes)
}
