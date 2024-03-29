package controller

import (
	"encoding/json"
	"hackathon/model"
	"hackathon/usecase"
	"io"
	"log"
	"net/http"
)

type RegisterMessageController struct {
	RegisterMessageUseCase *usecase.RegisterMessageUseCase
}

func (c *RegisterMessageController) Handle(w http.ResponseWriter, r *http.Request) {
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
	var userReq model.MessageReqForHTTPPost
	if err := json.Unmarshal(body, &userReq); err != nil {
		log.Printf("fail: json.Unmarshal, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user, err := c.RegisterMessageUseCase.Handle(userReq)
	if err != nil {
		log.Printf("fail: RegisterMessageUseCase.Handle, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	userRes := model.MessageResForHTTPPost{Id: user.Id}
	bytes, err := json.Marshal(userRes)
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
