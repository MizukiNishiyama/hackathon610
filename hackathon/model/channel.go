package model

type Channel struct {
	Id   string
	Name string
}

type ChannelResForHTTPGet struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type ChannelResForHTTPPost struct {
	Id string `json:"id"`
}

type ChannelReqForHTTPPost struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
