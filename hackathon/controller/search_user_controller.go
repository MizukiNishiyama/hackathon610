package controller

import (
	"encoding/json"
	"hackathon/model"
	"hackathon/usecase"
	"log"
	"net/http"
)

type SearchUserController struct {
	SearchUserUseCase *usecase.SearchUserUseCase
}

func (c *SearchUserController) Handle(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		log.Println("fail: name is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	users, err := c.SearchUserUseCase.Handle(name)
	if err != nil {
		log.Printf("fail: SearchUserUseCase.Handle, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	usersRes := make([]model.UserResForHTTPGet, len(users))
	for i, u := range users {
		usersRes[i] = model.UserResForHTTPGet{Id: u.Id, Name: u.Name, Email: u.Email}
	}

	bytes, err := json.Marshal(usersRes)
	if err != nil {
		log.Printf("fail: json.Marshal, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	w.Write(bytes)
}
