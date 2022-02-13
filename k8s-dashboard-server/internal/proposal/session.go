package proposal

import (
	"encoding/json"
)

type SessionUserInfo struct {
	UserID   int32  `json:"user_id"`
	UserName string `json:"user_name"`
}

func (user *SessionUserInfo) Marshal() (jsonRaw []byte) {
	jsonRaw, _ = json.Marshal(user)
	return
}
