package config

// AppResponse is for response config to Frontend side
type AppResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
