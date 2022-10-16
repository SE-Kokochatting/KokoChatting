package internal

type Message interface {
	Json() []byte
	MessageType() int
}
