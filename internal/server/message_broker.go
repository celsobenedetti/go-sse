package server

import "fmt"

type Message struct {
	Id       string
	RoomId   string
	SenderId string
	Message  string
}

type MessageBroker interface {
	Subscribe(roomId, userId string) (
		ch chan Message,
		unsub func() error,
		err error,
	)
	Publish(msg Message) error
	Health() (string, error)
}

var _ MessageBroker = (*InMemoryMessageBroker)(nil)

// WARN: definitely not thread safe
type InMemoryMessageBroker struct {
	// Map of [ roomId ]  => Subscribers
	// Each Subscriber is a map of [ userIds ] => chan Message
	rooms map[string]map[string]chan Message
}

func (b *InMemoryMessageBroker) Health() (string, error) {
	return "InMemoryMessageBroker is healthy", nil
}

func NewInMemoryMessageBroker() *InMemoryMessageBroker {
	return &InMemoryMessageBroker{
		rooms: make(map[string]map[string]chan Message),
	}
}

func (b *InMemoryMessageBroker) Subscribe(roomId, userId string) (
	ch chan Message,
	unsub func() error,
	err error,
) {
	unsub = func() error {
		fmt.Println("unsub called")
		delete(b.rooms[roomId], userId)
		return nil
	}

	_, ok := b.rooms[roomId]
	if !ok {
		b.rooms[roomId] = make(map[string]chan Message)
	}

	ch = make(chan Message)
	b.rooms[roomId][userId] = ch
	return ch, unsub, nil
}

func (b *InMemoryMessageBroker) Publish(msg Message) error {
	_, ok := b.rooms[msg.RoomId]
	if !ok {
		return fmt.Errorf("room does not exist")
	}
	fmt.Println(len(b.rooms[msg.RoomId]))

	subscribers := b.rooms[msg.RoomId]
	for _, ch := range subscribers {
		ch <- msg
	}
	return nil
}
