package api

import (
	"bytes"
	"chaincode/model"
	"chaincode/pkg/utils"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	"image/png"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// 序列化函数
func SerializeImage(img image.Image) (string, error) {
	buf := new(bytes.Buffer)
	err := png.Encode(buf, img) // 将图像保存为PNG格式字节切片
	if err != nil {
		return "", err
	}
	encoded := base64.StdEncoding.EncodeToString(buf.Bytes()) // 对字节切片进行base64编码
	return encoded, nil
}

// 反序列化函数
func DeserializeImage(encoded string) (image.Image, error) {
	decoded, err := base64.StdEncoding.DecodeString(encoded) // 对base64编码后的字符串进行解码
	if err != nil {
		return nil, err
	}
	img, err := png.Decode(bytes.NewBuffer(decoded)) // 将解码后的字节切片解析为图像
	if err != nil {
		return nil, err
	}
	return img, nil
}

func CreatePicture(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// 验证参数
	fmt.Println(args)
	if len(args) < 4 {
		return shim.Error("参数个数不满足")
	}
	publisher := args[0] //发布者|权限
	time := args[1]      //时间
	image := args[2]
	names := args[3]

	picture := &model.Picture{
		AccountID: publisher, //发布者|权限
		Time:      time,      //时间
		Image:     image,     //图像
		Names:     names,     //拍摄照片内包含的人名列表
	}
	// 写入账本
	if err := utils.WriteLedger(picture, stub, model.PictureKey, []string{picture.AccountID, picture.Time}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	//将成功创建的信息返回
	pictureByte, err := json.Marshal(picture)
	if err != nil {
		return shim.Error(fmt.Sprintf("序列化成功创建的信息出错: %s", err))
	}
	// 成功返回
	return shim.Success(pictureByte)
}

// 查询图像记录
func QueryPicture(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var realPictureList []interface{}
	//var realPictureList []model.Picture
	results, err := utils.GetStateByPartialCompositeKeys(stub, model.PictureKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}

	for _, v := range results {
		if v != nil {
			var realPicture interface{}
			//var realPicture model.Picture
			err := json.Unmarshal(v, &realPicture)
			if err != nil {
				return shim.Error(fmt.Sprintf("QueryPicture-反序列化出错:%s", err))
			}
			realPictureList = append(realPictureList, realPicture)
		}
	}
	fmt.Sprintln(realPictureList...)
	realPictureListByte, err := json.Marshal(realPictureList)
	if err != nil {
		shim.Error(fmt.Sprintf("QueryPicture-序列化出错:%s", err))
	}
	return shim.Success(realPictureListByte)
}
