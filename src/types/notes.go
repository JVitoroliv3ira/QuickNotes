package types

type Notes struct {
	Data      map[string]Note `json:"data"`
	CurrentId int             `json:"current_id"`
}
