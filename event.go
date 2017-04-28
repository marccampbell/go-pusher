package pusher

import "fmt"

// EventStub contains just the "type" of event.
// Knowing the type, we can then unmarshal again, into appropriate type.
type EventStub struct {
	Event string `json:"event"`
}

type Event struct {
	Event   string      `json:"event"`
	Data    interface{} `json:"data"`
	Channel string      `json:"channel"`
}

// EventError contains a structured error in its Data field.
// It implements error.
type EventError struct {
	Event string `json:"event"`
	Data  struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
	} `json:"data"`
}

type ConnectionData struct {
	ActivityTimeout int    `json:"activity_timeout"`
	SocketID        string `json:"socket_id"`
}

type AuthResponse struct {
	Auth string `json:"auth"`
}

func (ewe EventError) Error() string {
	return fmt.Sprintf("Pusher return error: code %d, message %q", ewe.Data.Code, ewe.Data.Message)
}
