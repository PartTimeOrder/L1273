package model

// Result 用于接收 json 中的 result
type Result struct {
	Hash             string `json:"hash"`
	TransactionIndex string `json:"transactionIndex"`
	From             string `json:"from"`
	To               string `json:"to"`
	Value            string `json:"value"`
	Input            string `json:"input"`
	TimeStamp        string `json:"timeStamp"`
}

// Response 用于接收返回的 json
type Response struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Result  []Result `json:"result"`
}
