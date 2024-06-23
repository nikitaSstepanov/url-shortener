package entity

type Message struct {
	Msg    string
	Status StatusEnum
}

const (
	NotFound StatusEnum = iota
	BadInput
	Internal
	Conflict
)

type StatusEnum int

func GetMsg(msg string, status StatusEnum) *Message {
	return &Message{
		Msg:    msg,
		Status: status,
	}
}