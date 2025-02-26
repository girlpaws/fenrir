package fenrir

type Webhook struct {
	ID        string      `json:"id"`
	Name      string      `json:"name"`
	Avatar    *Attachment `json:"avatar"`
	ChannelID string      `json:"channel_id"`
	Token     string      `json:"token"`
}

type WebhookCreate struct {
	Name   string `json:"name"`
	Avatar string `json:"avatar,omitempty"`
}
