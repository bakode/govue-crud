package entities

type ResponseMany struct {
	Status  int       `json:"status"`
	Message string    `json:"message"`
	Data    []Example `json:"results"`
}

type ResponseRow struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Data    Example `json:"result"`
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
