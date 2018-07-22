package message

// MessageListener recieves messages from a message handler for processing
type MessageListener interface {
	MessageReceived(m Message)
}
