package models

type ApiResult struct {
	Source     string      `json:"source"`
	Data       interface{} `json:"data"`
	Priorities map[string]int
}

type Job struct {
	URL        string
	Priorities map[string]int
}
