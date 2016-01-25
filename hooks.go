package gomm

/* Types taken from mattermost source code */

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type OutgoingWebhook struct {
	Id        string `json:"id"`
	Token     string `json:"token"`
	CreateAt  int64  `json:"create_at"`
	UpdateAt  int64  `json:"update_at"`
	DeleteAt  int64  `json:"delete_at"`
	CreatorId string `json:"creator_id"`
	ChannelId string `json:"channel_id"`
	TeamId    string `json:"team_id"`
	//TriggerWords StringArray `json:"trigger_words"`
	//CallbackURLs StringArray `json:"callback_urls"`
}

type IncomingWebhookRequest struct {
	Text        string `json:"text"`
	Username    string `json:"username"`
	IconURL     string `json:"icon_url"`
	ChannelName string `json:"channel"`
	//Props       StringInterface `json:"props"`
	Attachments interface{} `json:"attachments"`
	Type        string      `json:"type"`
}

type Mattermost struct {
	IncomingURL string
}

func (m *Mattermost) Send(req IncomingWebhookRequest) error {
	client := http.Client{}

	enc, err := json.Marshal(req)
	if err != nil {
		return err
	}

	client.Post(m.IncomingURL, "application/json", bytes.NewBuffer(enc))
	return nil
}
