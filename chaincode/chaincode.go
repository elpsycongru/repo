package main

import (
	"chaincode/api"
	"chaincode/model"
	"chaincode/pkg/utils"
	"fmt"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type BlockChainRealEstate struct {
}

// Init 链码初始化
func (t *BlockChainRealEstate) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("链码初始化")
	//初始化默认数据
	var accountIds = [3]string{
		"5feceb66ffc8",
		"6b86b273ff34",
		"d4735e3a265e",
	}
	var userNames = [3]string{"Peer0", "Peer1", "Peer2"}
	//初始化账号数据
	for i, val := range accountIds {
		account := &model.Account{
			AccountID: val,
			UserName:  userNames[i],
			Balance:   0,
		}
		// 写入账本
		//参数含义：{传入数据的类型，接口，记录类型，主键列表}
		if err := utils.WriteLedger(account, stub, model.AccountKey, []string{val}); err != nil {
			return shim.Error(fmt.Sprintf("%s", err))
		}
	}
	return shim.Success(nil)
}

// Invoke 实现Invoke接口调用智能合约
func (t *BlockChainRealEstate) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	funcName, args := stub.GetFunctionAndParameters()
	switch funcName {
	case "hello":
		return api.Hello(stub, args)
		// return api.CreateRealEstate(stub, args)
	case "createPicture":
		return api.CreatePicture(stub, args)
	case "queryPicture":
		return api.QueryPicture(stub, args)
	default:
		return shim.Error(fmt.Sprintf("没有该功能: %s", funcName))
	}
}

func main() {
	timeLocal, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
	time.Local = timeLocal
	err = shim.Start(new(BlockChainRealEstate))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
