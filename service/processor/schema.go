package processor

type Request struct {
	Feature []float64 `json:"feature"`
}

type baseResponse struct {
	ErrCode int    `json:"err_code"`
	Msg     string `json:"msg"`
}

type prediction struct {
	Temperature float64 `json:"temperature"`
}

type Response struct {
	*baseResponse
	*prediction
}
