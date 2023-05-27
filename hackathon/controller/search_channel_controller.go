package controller

import (
	"encoding/json"
	"hackathon/model"
	"hackathon/usecase"
	"log"
	"net/http"
)

type SearchChannelController struct {
	SearchChannelUseCase *usecase.SearchChannelUseCase
}

func (c *SearchChannelController) Handle(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		log.Println("fail: name is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	channels, err := c.SearchChannelUseCase.Handle(name)
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
	w.Write(bytes)
}
