package telegram

type Update struct {
	ID      int             `json:"update_id"`
	Message *Incomingessage `json:"message"`
}

type UpdatesResponse struct {
	Ok     bool     `json:"ok"`
	Result []Update `json:"result"`
}

type Incomingessage struct {
	Text string `json:"text"`
	From From   `json:"from"`
	Chat Chat   `json:"chat"`
}

type Chat struct {
	ID int `json:"id"`
}

type From struct {
	Username string `json:"username"`
}
