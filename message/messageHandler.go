package message

// MessageHandler distributes messages to message listeners
type MessageHandler struct {
	m         Message
	listeners map[MessageListener]bool
}

// NewMessageHandler creates a new instance of the MessageHandler struct
func NewMessageHandler() *MessageHandler {
	return &MessageHandler{
		listeners: map[MessageListener]bool{},
	}
}

// AddListener inserts the provided listener in the handler's collection of
// listeners.  The collection will include each listener only once
func (mh *MessageHandler) AddListener(ml MessageListener) {
	mh.listeners[ml] = true
}

// RemoveListener removes the provided listener from the handler's collection
func (mh *MessageHandler) RemoveListener(ml MessageListener) {
	delete(mh.listeners, ml)
}

// SendMessage keeps the most recent message and notifies all connected message
// listeners of it
func (mh *MessageHandler) SendMessage(m Message) {
	mh.m = m
	for ml := range mh.listeners {
		ml.MessageReceived(m)
	}
}
