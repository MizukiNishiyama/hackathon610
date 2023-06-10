package controller

import (
	"hackathon/usecase"
	"log"
	"net/http"
)

type DeleteChannelController struct {
	DeleteChannelUseCase *usecase.DeleteChannelUseCase
}

func (c *DeleteChannelController) Handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "missing message id", http.StatusBadRequest)
	}
	err := c.DeleteChannelUseCase.DeleteChannel(id)
	if err != nil {
		log.Printf("fail: DeleteChannelUseCase.Handle, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
