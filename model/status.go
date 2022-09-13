package model

type Status struct {
	Online  bool  `json:"online"`
	Running bool  `json:"running"`
	Last    int64 `json:"last"` //time.Now().Unix()
}
