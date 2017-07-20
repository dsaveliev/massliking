package models

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type InstagramCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Trusted  bool   `json:"trusted"`
	Speed    `json:"speed"`
	Hours    `json:"hours"`
}

type ChannelCredentials struct {
	Value    string   `json:"value"`
	Action   string   `json:"action"`
	Target   string   `json:"target"`
	Comments []string `json:"comments"`
}
