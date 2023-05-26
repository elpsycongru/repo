package main

import (
	"bytes"
	"chaincode/api"
	"chaincode/model"
	"encoding/json"
	"fmt"
	"image/png"
	"os"
	"testing"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

func initTest(t *testing.T) *shim.MockStub {
	scc := new(BlockChainRealEstate)
	//"ex01"为名称
	stub := shim.NewMockStub("ex01", scc)
	checkInit(t, stub, [][]byte{[]byte("init")})
	return stub
}

func checkInit(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInit("1", args)
	if res.Status != shim.OK {
		fmt.Println("初始化失败", string(res.Message))
		t.FailNow()
	}
}

func checkInvoke(t *testing.T, stub *shim.MockStub, args [][]byte) pb.Response {
	//- 1为uuid，用于链码开始前和结束后开始事务的标志，无实际意义
	//- args为初始化需要的参数
	res := stub.MockInvoke("1", args)
	if res.Status != shim.OK {
		fmt.Println("Invoke", args, "failed", string(res.Message))
		t.FailNow()
	}
	return res
}

func TestPictureInit(t *testing.T) {
	initTest(t)
}

// 测试链码初始化

func TestHello(t *testing.T) {
	stub := initTest(t)
	fmt.Println(fmt.Sprintf("1、测试hello功能\n%s",
		string(checkInvoke(t, stub, [][]byte{
			[]byte("hello"),
		}).Payload)))
}
func getImage() string {
	file, err := os.Open("./facerecognition/ramdisk/test.png")
	if err != nil {
		fmt.Println("图像打开失败")
		return ""
	}
	img, err := png.Decode(file)
	if err != nil {
		fmt.Println("图像打开为png格式失败")
		return ""
	}
	msg, err := api.SerializeImage(img)
	if err != nil {
		fmt.Println("图像序化失败")
		return ""
	}

	return msg
}

// 返回当前时间
func getTimeNow() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

/*
	func TestCreatePicture(t *testing.T) {
		stub := initTest(t)
		time := getTimeNow()
		image := getImage()
		names := facerecognition.Face_detect()
		fmt.Println("\n", fmt.Sprintf("2、测试创建图像信息功能\n%s",
			string(checkInvoke(t, stub, [][]byte{
				[]byte("createPicture"),
				[]byte("Peer0"),
				[]byte(time),
				[]byte(image),
				[]byte(names),
			}).Payload)))
		fmt.Print("3、测试查询图像信息功能\n")
		_ = checkInvoke(t, stub, [][]byte{
			[]byte("queryPicture"),
		}).Payload
	}
*/
func TestQueryPcture(t *testing.T) {
	stub := initTest(t)
	_ = checkCreatepicture(stub, t)

	fmt.Println(fmt.Sprintf("3、测试查询图像信息功能\n%s",
		string(checkInvoke(t, stub, [][]byte{
			[]byte("queryPicture"),
		}).Payload)))
}

// 手动创建一些图像记录
func checkCreatepicture(stub *shim.MockStub, t *testing.T) []model.Picture {
	var pictureList []model.Picture
	var picture model.Picture
	//手动创建记录调用链码
	resp1 := checkInvoke(t, stub, [][]byte{
		[]byte("createPicture"),
		[]byte("Peer0"),
		[]byte("2023-05-16 11:48:06"),
		[]byte(""),
		[]byte(""),
	})

	resp2 := checkInvoke(t, stub, [][]byte{
		[]byte("createPicture"),
		[]byte("Peer1"),
		[]byte("2023-05-16 12:35:15"),
		[]byte(""),
		[]byte(""),
	})

	resp3 := checkInvoke(t, stub, [][]byte{
		[]byte("createPicture"),
		[]byte("Peer2"),
		[]byte("2023-05-16 13:22:50"),
		[]byte(""),
		[]byte(""),
	})

	resp4 := checkInvoke(t, stub, [][]byte{
		[]byte("createPicture"),
		[]byte("Peer0"),
		[]byte("2023-05-16 14:05:32"),
		[]byte(""),
		[]byte(""),
	})
	//将创建的记录保存起来返回
	json.Unmarshal(bytes.NewBuffer(resp1.Payload).Bytes(), &picture)
	pictureList = append(pictureList, picture)
	json.Unmarshal(bytes.NewBuffer(resp2.Payload).Bytes(), &picture)
	pictureList = append(pictureList, picture)
	json.Unmarshal(bytes.NewBuffer(resp3.Payload).Bytes(), &picture)
	pictureList = append(pictureList, picture)
	json.Unmarshal(bytes.NewBuffer(resp4.Payload).Bytes(), &picture)
	pictureList = append(pictureList, picture)
	return pictureList
}

/*
func TestFaceRecognition(t *testing.T) {
	names := facerecognition.Face_detect()
	fmt.Println(names)
}
*/
