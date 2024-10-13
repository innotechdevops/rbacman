package core

type Params struct {
	Offset int64 `json:"offset"`
	Page   int64 `json:"page"`
	Limit  int64 `json:"limit"`
}
