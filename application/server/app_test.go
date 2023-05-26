package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestHello(t *testing.T) {
	url := "http://localhost:8080/ping"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("请求失败：", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应失败：", err)
		return
	}

	fmt.Println(string(body))
}
