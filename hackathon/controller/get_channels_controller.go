package controller

import (
	"encoding/json"
	"hackathon/model"
	"hackathon/usecase"
	"log"
	"net/http"
)

type GetChannelController struct {
	GetChannelUseCase *usecase.GetChannelUseCase
}

func (c *GetChannelController) Handle(w http.ResponseWriter, r *http.Request) {

	channels, err := c.GetChannelUseCase.Handle()
	if err != nil {
		log.Printf("fail: SearchChannelUseCase.Handle, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	channelsRes := make([]model.ChannelResForHTTPGet, len(channels))
	for i, u := range channels {
		channelsRes[i] = model.ChannelResForHTTPGet{Id: u.Id, Name: u.Name}
	}

	bytes, err := json.Marshal(channelsRes)
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
