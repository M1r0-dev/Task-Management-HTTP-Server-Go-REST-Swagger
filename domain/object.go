package domain

type Task struct {
	ID       string `json:"id"`
	Code     string `json:"code"`
	Compiler string `json:"compiler"`
	Status   string `json:"status"`
	Result   string `json:"result"`
}
