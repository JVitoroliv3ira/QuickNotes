package types

type Notes struct {
	Data map[string]struct {
		Id   string `json:"id"`
		Text string `json:"text"`
	} `json:"data"`
}
