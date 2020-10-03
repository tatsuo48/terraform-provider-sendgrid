package sendgrid

type whitelistResponse struct {
	Result []struct {
		ID        int    `json:"id"`
		IP        string `json:"ip"`
		CreatedAt int    `json:"created_at"`
		UpdatedAt int    `json:"updated_at"`
	} `json:"result"`
}

type whitelistIPResponse struct {
	Result struct {
		CreatedAt int    `json:"created_at"`
		ID        int    `json:"id"`
		IP        string `json:"ip"`
		UpdatedAt int    `json:"updated_at"`
	} `json:"result"`
}
