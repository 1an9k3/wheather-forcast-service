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
	// get json req
	req := new(Request)
	bytesJSONReq, _ := ctx.GetRawData()
	err := json.Unmarshal(bytesJSONReq, req)
	if err != nil {
		log.Panicf("unmarshal fail:%v", err)
	}
	// type validation
	inputFea := req.feature
	switch len(inputFea) {
	}
}
