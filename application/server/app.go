package main

import (
	"fmt"
	"log"
	"net/http"

	bc "application/blockchain"
	"application/routers"
)

func main() {
	/*
		timeLocal, err := time.LoadLocation("Asia/Shanghai")
		if err != nil {
			log.Printf("时区设置失败 %s", err)
		}
		time.Local = timeLocal*/
	bc.Init()

	endPoint := fmt.Sprintf("0.0.0.0:%d", 8000)
	server := &http.Server{
		Addr:    endPoint,
		Handler: routers.InitRouter(),
	}
	log.Printf("[info] start http server listening %s", endPoint)
	if err := server.ListenAndServe(); err != nil {
		log.Printf("start http server failed %s", err)
	}

}
