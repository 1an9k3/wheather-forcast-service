package processor

type Request struct {
	Feature []float32 `json:"feature"`
}

type baseResponse struct {
	ErrCode int    `json:"err_code"`
	Msg     string `json:"msg"`
}

type prediction struct {
	Temperature float32 `json:"temperature"`
}

type Response struct {
	*baseResponse
	*prediction
}
