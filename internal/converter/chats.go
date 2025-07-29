package converter

import (
	"github.com/ldevprog/chat-server/internal/model"
	desc "github.com/ldevprog/chat-server/pkg/chat_v1"
)

func ToChatCreateFromDesc(req *desc.CreateRequest) *model.ChatCreate {
	return &model.ChatCreate{
		UserNames: req.GetUsernames(),
	}
}

func ToMessageCreateFromDesct(req *desc.SendMessageRequest) *model.MessageCreate {
	return &model.MessageCreate{
		ChatId:    req.GetChatId(),
		From:      req.GetFrom(),
		Text:      req.GetText(),
		Timestamp: req.Timestamp.AsTime(),
	}
}
