package server

import "encoding/json"

type Message struct {
	Id       string `json:"id,omitempty"`
	RoomId   string `json:"roomId,omitempty"`
	SenderId string `json:"senderId,omitempty"`
	Message  string `json:"message,omitempty"`
}

func (m Message) MarshalBinary() (data []byte, err error) {
	return json.Marshal(m)
}
