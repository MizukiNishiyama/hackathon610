package controller

import (
	"hackathon/usecase"
	"log"
	"net/http"
)

type DeleteMessageController struct {
	DeleteMessageUseCase *usecase.DeleteMessageUseCase
}

func (c *DeleteMessageController) Handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	messageID := r.URL.Query().Get("id")
	if messageID == "" {
		http.Error(w, "missing message id", http.StatusBadRequest)
	}
	err := c.DeleteMessageUseCase.DeleteMessage(messageID)
	if err != nil {
		log.Printf("fail: DeleteMessageUseCase.Handle, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
