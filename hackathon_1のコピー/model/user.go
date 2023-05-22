package model

type User struct {
	Id   string
	Name string
	Age  int
}

type UserResForHTTPGet struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type UserResForHTTPPost struct {
	Id string `json:"id"`
}

type UserReqForHTTPPost struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
