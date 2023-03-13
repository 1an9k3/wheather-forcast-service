package service

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, BaseResponse{
		ErrCode: ERROR_OK,
	})
}

func predict(ctx *gin.Context) {
	var resp Response
	var errCode int
	// get global config
	tmpCfg, ok := ctx.Get("config")
	if !ok {
		log.Println("load config error")
		resp.ErrCode = ERROR_LOAD_CONFIG
		errCode = http.StatusInternalServerError
	}
	gc := tmpCfg.(*TomlConfig)

	// get json req
	req := new(Request)
	bytesJSONReq, _ := ctx.GetRawData()
	err := json.Unmarshal(bytesJSONReq, req)
	if err != nil {
		log.Printf("unmarshal fail:%v", err)
		errCode = ERROR_LOAD_CONFIG
		errCode = http.StatusBadRequest
	}
	// type validation
	inputFea := req.feature

	if len(inputFea) != gc.Predictor.FeatureLen {

	}
}
