package api

import (
	bc "application/blockchain"
	"application/pkg/app"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Response struct {
	Status string `json:"message"`
	Names  string `json:"names"`
}
type config struct {
	Pyport     string `json:"python_server_port"`
	Serverport string `json:"serverport"`
	Imgdir     string `json:"unknown_img_dir"`
}

var (
	Config config
)

// 返回当前时间
func getTimeNow() []byte {
	return []byte(time.Now().Format("2006-01-02 15:04:05"))
}

func CreatePicture(c *gin.Context) {

	appG := app.Gin{C: c}
	var image []byte
	//识别图像中的人脸信息
	file, err := c.FormFile("img")
	m := map[string]string{"name": "empty"}
	var namelist string
	if err != nil {
		m["name"] = "error"
		appG.Response(http.StatusBadRequest, "传输图像为空", err)
		return
	} else {
		//data := make([]byte, file.Size)
		f, err := file.Open()
		if err != nil {
			appG.Response(http.StatusBadRequest, "图像打开失败", err)
		}
		defer f.Close()
		data, err := ioutil.ReadAll(f)
		image = data
		//
		name := Config.Imgdir + "/" + uuid.NewString() + path.Ext(file.Filename)
		os.WriteFile(name, data, 0666)
		namelist, err = getNames(name)
		if err != nil {
			fmt.Println(err)
			m["name"] = err.Error()
		} else {
			m["name"] = namelist
		}
		fmt.Print(m["name"])
		os.Remove(name)

	}

	accountId := c.DefaultQuery("account", "Peer0")
	time := getTimeNow()

	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(accountId))
	bodyBytes = append(bodyBytes, time)
	bodyBytes = append(bodyBytes, image)
	bodyBytes = append(bodyBytes, []byte(namelist))
	// 调用智能合约写入账本
	resp, err := bc.ChannelExecute("createPicture", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}

	var data map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}

// 查询所有图像记录
func QueryPicture(c *gin.Context) {
	/*
		appG := app.Gin{C: c}
		data := []map[string]interface{}{
			{"AccountID": "Peer0", "Time": "2023-05-20 09:30:35", "Names": "Trump,Biden"},
			{"AccountID": "Peer2", "Time": "2023-05-20 09:46:28", "Names": "Obama"},
			{"AccountID": "Peer1", "Time": "2023-05-20 10:01:37", "Names": "Trump"},
			{"AccountID": "Peer1", "Time": "2023-05-20 10:05:55", "Names": "Biden"},
			{"AccountID": "Peer0", "Time": "2023-05-20 10:13:42", "Names": "Obama"},
			{"AccountID": "Peer1", "Time": "2023-05-20 10:25:16", "Names": "Biden"},
			{"AccountID": "Peer2", "Time": "2023-05-20 11:07:20", "Names": "Biden"},
		}
		appG.Response(http.StatusOK, "成功", data)
		return
	*/

	appG := app.Gin{C: c}
	//解析body参数,暂无

	bodyBytes := [][]byte{}
	// 调用智能合约，查询
	resp, err := bc.ChannelQuery("queryPicture", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}

	var data map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)

}

func QueryByName(c *gin.Context) {
	appG := app.Gin{C: c}
	data := []map[string]interface{}{
		{"AccountID": "Peer0", "Time": "2023-05-20 09:30:35", "Names": "Trump,Biden"},
		{"AccountID": "Peer1", "Time": "2023-05-20 10:05:55", "Names": "Biden"},
		{"AccountID": "Peer1", "Time": "2023-05-20 10:25:16", "Names": "Biden"},
		{"AccountID": "Peer2", "Time": "2023-05-20 11:07:20", "Names": "Biden"},
	}
	appG.Response(http.StatusOK, "成功", data)
}
func QueryByPeer(c *gin.Context) {
	appG := app.Gin{C: c}
	data := []map[string]interface{}{
		{"AccountID": "Peer1", "Time": "2023-05-20 10:01:37", "Names": "Trump"},
		{"AccountID": "Peer1", "Time": "2023-05-20 10:05:55", "Names": "Biden"},
		{"AccountID": "Peer1", "Time": "2023-05-20 10:25:16", "Names": "Biden"},
	}
	appG.Response(http.StatusOK, "成功", data)
}

func QueryBetweenTime(c *gin.Context) {
	appG := app.Gin{C: c}
	data := []map[string]interface{}{
		{"AccountID": "Peer1", "Time": "2023-05-20 10:01:37", "Names": "Trump"},
		{"AccountID": "Peer1", "Time": "2023-05-20 10:05:55", "Names": "Biden"},
		{"AccountID": "Peer0", "Time": "2023-05-20 10:13:42", "Names": "Obama"},
		{"AccountID": "Peer1", "Time": "2023-05-20 10:25:16", "Names": "Biden"},
	}
	appG.Response(http.StatusOK, "成功", data)
}

func getNames(filename string) (string, error) {
	url := "http://0.0.0.0:43221"
	//url := "http://face_recognition.app:43221/"

	jsonStr := []byte(`{"path": "test.png"}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return "", err
	}
	req.Header.Set("Cotent-Type", "applicaiton/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	fmt.Println(resp.Body)
	defer resp.Body.Close()

	var response Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return "", err
	}
	fmt.Println(response)
	str := response.Names
	fmt.Println(str)
	return str, nil

}
func init() {
	//读取配置文件
	/*
		configfile, err := os.ReadFile("/config/config.json")
		if err != nil {
			log.Fatal(err)
		}
		//解码
		err = json.Unmarshal(configfile, &Config)
		if err != nil {
			log.Fatal(err)
		}
	*/

	Config.Pyport = ":" + "43221"
	Config.Serverport = ":" + "9090"
	Config.Imgdir = "/Users/lichenhui/go/src/github.com/elpsycongru/communication-system-blockchain/tools/facerecognition/ramdisk"

	//初始化缓存redis
	/*
		fmt.Println("Listening at post " + Config.Serverport)
		http.HandleFunc("/", face)
		log.Fatal(http.ListenAndServe(Config.Serverport, nil))
	*/
}
