package server

import "encoding/json"

type Message struct {
	Id       string `json:"id"`
	RoomId   string `json:"roomId"`
	SenderId string `json:"senderId"`
	Message  string `json:"message"`
}

func (m Message) MarshalBinary() (data []byte, err error) {
	return json.Marshal(m)
}
