package payloads

type AuthPayload struct {
	Key string `json:"key"`
}

type AuthMessage struct {
	Message string `json"message"`
}

type CheckPayload struct {
	Message string `json:"message"`
}
