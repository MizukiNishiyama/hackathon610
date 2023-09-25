package controller

import (
	"fmt"
	"hackathon/usecase"
	"log"
	"net/http"
)

type EditMessageController struct {
	EditMessageUseCase *usecase.EditMessageUseCase
}

func (c *EditMessageController) Handle(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	content := r.URL.Query().Get("content")
	if id == "" {
		http.Error(w, "missing message id", http.StatusBadRequest)
	}
	fmt.Println(id, content)
	err := c.EditMessageUseCase.EditMessage(id, content)
	if err != nil {
		log.Printf("fail: EditMessageUseCase.Handle, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
