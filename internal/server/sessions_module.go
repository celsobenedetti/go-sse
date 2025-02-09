package server

type SessionsModule struct {
	store *SessionsStore
}

func NewSessionsModule(mongo *MongoClient) *SessionsModule {
	return &SessionsModule{
		store: NewSessionsStore(mongo),
	}
}
