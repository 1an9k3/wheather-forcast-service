package processor

import (
	"encoding/json"
	"github.com/fyk1234/wheather-forcast-service/service/config"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Health(ctx *gin.Context) {
	var resp *Response
	resp.ErrCode = ERROR_OK
	ctx.JSON(http.StatusOK, resp)
}

func Predict(ctx *gin.Context) {
	var statusCode = http.StatusOK
	var errCode = 0
	// get global config
	tmpCfg, ok := ctx.Get("config")
	if !ok {
		log.Println("load config error")
		errCode = ERROR_LOAD_CONFIG
		statusCode = http.StatusInternalServerError
	}
	tomlCfg := tmpCfg.(*config.TomlConfig)

	// get json req
	req := new(Request)
	bytesJSONReq, _ := ctx.GetRawData()
	err := json.Unmarshal(bytesJSONReq, req)
	if err != nil {
		log.Printf("unmarshal fail:%v", err)
		errCode = ERROR_PARAM
		statusCode = http.StatusBadRequest
	}
	// length validation
	inputFea := req.Feature
	if len(inputFea) != tomlCfg.Predictor.FeatureLen {
		log.Printf("error length: %v", len(inputFea))
		errCode = ERROR_PARAM
		statusCode = http.StatusBadRequest
	}

	// TODO Predict
	var temperature float32
	// return
	var resp Response
	resp.ErrCode = errCode
	if statusCode == http.StatusOK {
		resp.Temperature = temperature
	}
	ctx.JSON(statusCode, resp)
}
