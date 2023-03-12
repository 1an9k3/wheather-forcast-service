package service

type Request struct {
	feature []float32
}

type BaseResponse struct {
	ErrCode int `json:"err_code"`
}

type Response struct {
	BaseResponse
	Temperature float32 `json:"temperature"`
}

const (
	ERROR_OK         = 0
	ERROR_PARAM      = 1
	ERROR_LOAD_MODEL = 2
	ERROR_INFERENCE  = 3
)
