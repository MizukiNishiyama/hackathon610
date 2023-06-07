package controller

import (
	"hackathon/usecase"
	"log"
	"net/http"
)

//type EditMessageController struct {
//	EditMessageUseCase *usecase.DeleteMessageUseCase
//}
//
//func (c *EditMessageController) Handle(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	w.Header().Set("Access-Control-Allow-Origin", "*")
//	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
//	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
//
//	var m model.EditMessage
//	err := json.NewDecoder(r.Body).Decode(&m)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//
//	err = uc.EditMessage(m.Id, m.Content)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//	id := r.URL.Query().Get("id")
//	if id == "" {
//		http.Error(w, "missing message id", http.StatusBadRequest)
//	}
//	err := c.EditMessageUseCase.DeleteMessage(id)
//	if err != nil {
//		log.Printf("fail: EditMessageUseCase.Handle, %v\n", err)
//		w.WriteHeader(http.StatusInternalServerError)
//		return
//	}
//
//	w.WriteHeader(http.StatusOK)
//}

type EditMessageController struct {
	EditMessageUseCase *usecase.EditMessageUseCase
}

func (c *EditMessageController) Handle(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	content := r.URL.Query().Get("content")
	if id == "" {
		http.Error(w, "missing message id", http.StatusBadRequest)
	}
	err := c.EditMessageUseCase.EditMessage(id, content)
	if err != nil {
		log.Printf("fail: EditMessageUseCase.Handle, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}