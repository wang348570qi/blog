package main

import (
	"fmt"
	"log"
	"net/http"
	"blog/models"
	"blog/pkg/gredis"
	"blog/pkg/logging"

	"blog/pkg/setting"
	"blog/routers"

	"blog/pkg/util"

	"github.com/gin-gonic/gin"
)

func init() {
	setting.Setup()
	models.Setup()
	logging.Setup()
	gredis.Setup()
	util.Setup()
}

func main() {
	gin.SetMode(setting.ServerSetting.RunMode)
	routersInit := routers.InitRouter()
	readTimeOut := setting.ServerSetting.ReadTimeout
	writeTimeOut := setting.ServerSetting.WriteTimeout
	maxHeaderBytes := 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)

	s := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeOut,
		WriteTimeout:   writeTimeOut,
		MaxHeaderBytes: maxHeaderBytes,
	}
	log.Printf("[info] start http server listening %s", endPoint)
	err := s.ListenAndServe()
	if err != nil {
		log.Printf("server err:%v", err)
	}

}
