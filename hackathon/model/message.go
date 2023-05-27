package model

type Message struct {
	Id   string
	Content string
	UserId  string
	ChannelId string
}

type MessageResForHTTPGet struct {
	Id   string `json:"id"`
	Content string `json:"content"`
	UserId  string    `json:"userid"`
	ChannelId string `json:"channelid"`
}

type MessageResForHTTPPost struct {
	Id string `json:"id"`
}

type MessageReqForHTTPPost struct {
	Id   string `json:"id"`
	Content string `json:"content"`
	UserId  string    `json:"userid"`
	ChannelId string `json:"channelid"`
}



