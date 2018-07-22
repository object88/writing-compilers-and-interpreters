package message

import "fmt"

type Message interface {
	GetMessageId() int
}

type BaseMessage struct {
	fmt.Stringer
	messageId int
}

func NewBaseMessage(messageId int) BaseMessage {
	return BaseMessage{
		messageId: messageId,
	}
}

func (bm BaseMessage) GetMessageId() int {
	return bm.messageId
}
