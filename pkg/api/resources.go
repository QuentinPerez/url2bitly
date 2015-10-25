package api

type Answer struct {
	StatusCode string      `json:"status_code"`
	StatusTxt  string      `json:"status_txt"`
	Data       interface{} `json:"data"`
}
