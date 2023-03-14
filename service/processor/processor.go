package processor

import (
	"encoding/json"
	"fmt"
	"github.com/fyk1234/wheather-forcast-service/service/config"
	"github.com/fyk1234/wheather-forcast-service/service/predictor"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Health(ctx *gin.Context) {
	var statusCode = http.StatusOK
	var errCode = ERROR_OK
	var msg string
	temp, ok := ctx.Get("config")
	tomlCfg := temp.(*config.TomlConfig)
	if !ok {
		msg = "load config error"
		log.Printf(msg)
		statusCode = http.StatusInternalServerError
		errCode = ERROR_LOAD_CONFIG
	} else {
		msg = fmt.Sprintf("Weather Forcast Service [%v]", tomlCfg.Version)
	}
	ctx.JSON(statusCode, Response{
		baseResponse: &baseResponse{
			ErrCode: errCode,
			Msg:     msg,
		},
		prediction: nil,
	})

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
	model, err := predictor.LoadModel()
	if err != nil {
		errCode = ERROR_LOAD_MODEL
		statusCode = http.StatusInternalServerError
	}
	temperature, _ := predictor.Predict(model, inputFea)
	// return
	var resp Response
	resp.ErrCode = errCode
	if statusCode == http.StatusOK {
		resp.Temperature = temperature
	}
	ctx.JSON(statusCode, resp)
}
