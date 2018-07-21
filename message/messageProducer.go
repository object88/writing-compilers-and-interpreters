package message

type MessageProducer interface {
	AddMessageListener(ml MessageListener)
	RemoveMessageListner(ml MessageListener)
	SendMessage(m Message)
}
