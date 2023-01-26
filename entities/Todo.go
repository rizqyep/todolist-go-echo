package entities

type CreateTodoRequest struct {
	Task string `json:"task"`
}

type UpdateTodoPayload struct {
	ID   int64
	Task string `json:"task,omitempty"`
}

type UpdateTodoRequest struct {
	Task string `json:"task,omitempty"`
}
