package message

type Message interface {
	GetMessageID() int
}

type BaseMessage struct {
	messageId int
}

func NewBaseMessage(messageId int) BaseMessage {
	return BaseMessage{
		messageId: messageId,
	}
}

func (bm *BaseMessage) GetMessageID() int {
	return bm.messageId
}
