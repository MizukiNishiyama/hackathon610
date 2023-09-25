package model

type User struct {
	Id   string
	Name string
	Email  string
}

type UserResForHTTPGet struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Email  string    `json:"email"`
}

type UserResForHTTPPost struct {
	Id string `json:"id"`
}

type UserReqForHTTPPost struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}



